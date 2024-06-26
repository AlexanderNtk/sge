package keeper_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/testutil"

	"github.com/sge-network/sge/app/params"
	"github.com/sge-network/sge/testutil/sample"
	"github.com/sge-network/sge/testutil/simapp"
	sgetypes "github.com/sge-network/sge/types"
	betmodulekeeper "github.com/sge-network/sge/x/bet/keeper"
	bettypes "github.com/sge-network/sge/x/bet/types"
	housetypes "github.com/sge-network/sge/x/house/types"
	markettypes "github.com/sge-network/sge/x/market/types"
	"github.com/sge-network/sge/x/subaccount/types"
)

var (
	bettor1      = sample.NativeAccAddress()
	bettor1Funds = sdkmath.NewInt(10).Mul(micro)
)

func TestMsgServer(t *testing.T) {
	app, k, msgServer, ctx := setupMsgServerAndApp(t)

	parm := k.GetParams(ctx)
	parm.DepositEnabled = true
	k.SetParams(ctx, parm)

	subAccOwner := sample.NativeAccAddress()
	subAccFunder := sample.NativeAccAddress()

	// do subaccount creation
	require.NoError(
		t,
		testutil.FundAccount(
			app.BankKeeper,
			ctx,
			subAccFunder,
			sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, subAccFunds)),
		),
	)

	_, err := msgServer.Create(sdk.WrapSDKContext(ctx), &types.MsgCreate{
		Creator: subAccFunder.String(),
		Owner:   subAccOwner.String(),
		LockedBalances: []types.LockedBalance{
			{
				UnlockTS: uint64(time.Now().Add(24 * time.Hour).Unix()),
				Amount:   subAccFunds,
			},
		},
	})
	require.NoError(t, err)

	// fund a bettor
	require.NoError(
		t,
		testutil.FundAccount(
			app.BankKeeper,
			ctx,
			bettor1,
			sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, subAccFunds)),
		),
	)

	// add market
	market := addTestMarket(t, app, ctx, false)

	// do house deposit
	deposit := sdkmath.NewInt(1000).Mul(micro)
	depResp, err := msgServer.HouseDeposit(sdk.WrapSDKContext(ctx), houseDepositMsg(t, subAccOwner, market.UID, deposit))
	require.NoError(t, err)
	// check spend
	subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
	require.True(t, exists)
	require.Equal(t, subBalance.SpentAmount, deposit)

	// place bet
	betMsgServer := betmodulekeeper.NewMsgServerImpl(*app.BetKeeper)
	_, err = betMsgServer.Wager(sdk.WrapSDKContext(ctx), testBetMsgWager(t, bettor1, bettor1Funds))
	require.NoError(t, err)

	participateFee := app.HouseKeeper.GetHouseParticipationFee(ctx).Mul(sdkmath.LegacyNewDecFromInt(deposit)).TruncateInt()
	bettorFee := sdkmath.NewInt(100)

	t.Run("house wins", func(t *testing.T) {
		ctx, _ := ctx.CacheContext()
		app.MarketKeeper.Resolve(ctx, *market, &markettypes.MarketResolutionTicketPayload{
			UID:            market.UID,
			ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
			WinnerOddsUIDs: []string{testOddsUID2},
			Status:         markettypes.MarketStatus_MARKET_STATUS_RESULT_DECLARED,
		})
		err := app.BetKeeper.BatchMarketSettlements(ctx)
		require.NoError(t, err)
		err = app.OrderbookKeeper.BatchOrderBookSettlements(ctx)
		require.NoError(t, err)

		subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
		require.True(t, exists)
		require.NoError(t, err)

		require.Equal(t, subBalance.SpentAmount.String(), sdkmath.ZeroInt().Add(participateFee).String())
		// check profits were forwarded to subacc owner
		ownerBalance := app.BankKeeper.GetAllBalances(ctx, subAccOwner)
		require.Equal(t,
			ownerBalance.AmountOf(params.DefaultBondDenom).String(),
			sdkmath.NewInt(10).Mul(micro).Sub(bettorFee).String())
	})

	t.Run("house loses", func(t *testing.T) {
		ctx, _ := ctx.CacheContext()
		app.MarketKeeper.Resolve(ctx, *market, &markettypes.MarketResolutionTicketPayload{
			UID:            market.UID,
			ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
			WinnerOddsUIDs: []string{testOddsUID1},
			Status:         markettypes.MarketStatus_MARKET_STATUS_RESULT_DECLARED,
		})
		err := app.BetKeeper.BatchMarketSettlements(ctx)
		require.NoError(t, err)
		err = app.OrderbookKeeper.BatchOrderBookSettlements(ctx)
		require.NoError(t, err)

		subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
		require.True(t, exists)
		require.NoError(t, err)

		require.Equal(t, subBalance.SpentAmount.String(), sdkmath.ZeroInt().Add(participateFee).String())
		require.Equal(t, subBalance.LostAmount, sdkmath.LegacyNewDecFromInt(bettor1Funds.Sub(bettorFee)).Mul(sdkmath.LegacyMustNewDecFromStr("3.2")).TruncateInt())
		// check profits were forwarded to subacc owner
		ownerBalance := app.BankKeeper.GetAllBalances(ctx, subAccOwner)
		require.Equal(t, ownerBalance.AmountOf(params.DefaultBondDenom), sdkmath.ZeroInt())
	})
	t.Run("house refund", func(t *testing.T) {
		ctx, _ := ctx.CacheContext()
		app.MarketKeeper.Resolve(ctx, *market, &markettypes.MarketResolutionTicketPayload{
			UID:            market.UID,
			ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
			WinnerOddsUIDs: []string{testOddsUID1},
			Status:         markettypes.MarketStatus_MARKET_STATUS_CANCELED,
		})
		err := app.BetKeeper.BatchMarketSettlements(ctx)
		require.NoError(t, err)
		err = app.OrderbookKeeper.BatchOrderBookSettlements(ctx)
		require.NoError(t, err)

		subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
		require.True(t, exists)
		require.NoError(t, err)

		require.Equal(t, subBalance.SpentAmount, sdkmath.ZeroInt())
		require.Equal(t, subBalance.LostAmount, sdkmath.ZeroInt())
		// check profits were forwarded to subacc owner
		ownerBalance := app.BankKeeper.GetAllBalances(ctx, subAccOwner)
		require.Equal(t, ownerBalance.AmountOf(params.DefaultBondDenom), sdkmath.ZeroInt())
	})

	// TODO: not participated in bet fulfillment.

	t.Run("withdrawal", func(t *testing.T) {
		ctx, _ := ctx.CacheContext()
		_, err := msgServer.HouseWithdraw(sdk.WrapSDKContext(ctx), &types.MsgHouseWithdraw{Msg: houseWithdrawMsg(t, subAccOwner, deposit, depResp.Response.ParticipationIndex)})
		require.NoError(t, err)

		// do subaccount balance check
		subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
		require.True(t, exists)

		require.Equal(t, subBalance.SpentAmount.String(), sdkmath.NewInt(131999680).String()) // NOTE: there was a match in the bet + participate fee
		require.Equal(t, subBalance.LostAmount.String(), sdkmath.ZeroInt().String())
	})

	t.Run("withdrawal and market refund with bet fulfillment", func(t *testing.T) {
		ctx, _ := ctx.CacheContext()

		_, err := msgServer.HouseWithdraw(sdk.WrapSDKContext(ctx), &types.MsgHouseWithdraw{Msg: houseWithdrawMsg(t, subAccOwner, deposit, depResp.Response.ParticipationIndex)})
		require.NoError(t, err)

		app.MarketKeeper.Resolve(ctx, *market, &markettypes.MarketResolutionTicketPayload{
			UID:            market.UID,
			ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
			WinnerOddsUIDs: []string{testOddsUID1},
			Status:         markettypes.MarketStatus_MARKET_STATUS_CANCELED,
		})
		err = app.BetKeeper.BatchMarketSettlements(ctx)
		require.NoError(t, err)
		err = app.OrderbookKeeper.BatchOrderBookSettlements(ctx)
		require.NoError(t, err)

		subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
		require.True(t, exists)
		require.NoError(t, err)

		require.Equal(t, subBalance.SpentAmount, sdkmath.ZeroInt())
		require.Equal(t, subBalance.LostAmount, sdkmath.ZeroInt())
		// check profits were forwarded to subacc owner
		ownerBalance := app.BankKeeper.GetAllBalances(ctx, subAccOwner)
		require.Equal(t, ownerBalance.AmountOf(params.DefaultBondDenom), sdkmath.ZeroInt())
	})
}

func TestHouseWithdrawal_MarketRefund(t *testing.T) {
	app, k, msgServer, ctx := setupMsgServerAndApp(t)

	parm := k.GetParams(ctx)
	parm.DepositEnabled = true
	k.SetParams(ctx, parm)

	subAccOwner := sample.NativeAccAddress()
	subAccFunder := sample.NativeAccAddress()
	// do subaccount creation
	require.NoError(
		t,
		testutil.FundAccount(
			app.BankKeeper,
			ctx,
			subAccFunder,
			sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, subAccFunds)),
		),
	)

	_, err := msgServer.Create(sdk.WrapSDKContext(ctx), &types.MsgCreate{
		Creator: subAccFunder.String(),
		Owner:   subAccOwner.String(),
		LockedBalances: []types.LockedBalance{
			{
				UnlockTS: uint64(time.Now().Add(24 * time.Hour).Unix()),
				Amount:   subAccFunds,
			},
		},
	})
	require.NoError(t, err)

	// fund a bettor
	require.NoError(
		t,
		testutil.FundAccount(
			app.BankKeeper,
			ctx,
			bettor1,
			sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, subAccFunds)),
		),
	)

	// add market
	market := addTestMarket(t, app, ctx, false)

	// do house deposit
	deposit := sdkmath.NewInt(1000).Mul(micro)
	depResp, err := msgServer.HouseDeposit(sdk.WrapSDKContext(ctx), houseDepositMsg(t, subAccOwner, market.UID, deposit))
	require.NoError(t, err)
	// check spend
	subBalance, exists := k.GetAccountSummary(ctx, subAccAddr)
	require.True(t, exists)
	require.Equal(t, subBalance.SpentAmount, deposit)

	// do house withdrawal
	_, err = msgServer.HouseWithdraw(sdk.WrapSDKContext(ctx), &types.MsgHouseWithdraw{Msg: houseWithdrawMsg(t, subAccOwner, deposit, depResp.Response.ParticipationIndex)})
	require.NoError(t, err)

	// we expect the balance to be the original one minus participation fee
	subBalance, exists = k.GetAccountSummary(ctx, subAccAddr)
	require.True(t, exists)
	require.Equal(t, subBalance.SpentAmount, sdkmath.NewInt(100).Mul(micro)) // all minus participation fee
	require.Equal(t, subBalance.LostAmount, sdkmath.ZeroInt())
	require.Equal(t, subBalance.DepositedAmount, subAccFunds)
	subBankBalance := app.BankKeeper.GetAllBalances(ctx, subAccAddr)
	require.Equal(t, subBankBalance.AmountOf(params.DefaultBondDenom), subAccFunds.Sub(sdkmath.NewInt(100).Mul(micro))) // original funds - fee

	// resolve market with refund
	app.MarketKeeper.Resolve(ctx, *market, &markettypes.MarketResolutionTicketPayload{
		UID:            market.UID,
		ResolutionTS:   uint64(ctx.BlockTime().Unix()) + 10000,
		WinnerOddsUIDs: []string{testOddsUID1},
		Status:         markettypes.MarketStatus_MARKET_STATUS_CANCELED,
	})
	err = app.BetKeeper.BatchMarketSettlements(ctx)
	require.NoError(t, err)
	err = app.OrderbookKeeper.BatchOrderBookSettlements(ctx)
	require.NoError(t, err)

	subBalance, exists = k.GetAccountSummary(ctx, subAccAddr)
	require.True(t, exists)
	require.NoError(t, err)

	require.Equal(t, subBalance.SpentAmount, sdkmath.ZeroInt())
	require.Equal(t, subBalance.LostAmount, sdkmath.ZeroInt())
	subBankBalance = app.BankKeeper.GetAllBalances(ctx, subAccAddr)
	require.Equal(t, subBankBalance.AmountOf(params.DefaultBondDenom), subAccFunds) // original funds - fee was refunded
	// check profits were not forwarded to subacc owner
	ownerBalance := app.BankKeeper.GetAllBalances(ctx, subAccOwner)
	require.Equal(t, ownerBalance.AmountOf(params.DefaultBondDenom), sdkmath.ZeroInt())
}

func houseWithdrawMsg(t testing.TB, owner sdk.AccAddress, amt sdkmath.Int, partecipationIndex uint64) *housetypes.MsgWithdraw {
	testKyc := &sgetypes.KycDataPayload{
		Approved: true,
		ID:       owner.String(),
	}
	ticketClaim := jwt.MapClaims{
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
		"iat":      time.Now().Unix(),
		"kyc_data": testKyc,
	}
	ticket, err := simapp.CreateJwtTicket(ticketClaim)
	require.Nil(t, err)

	inputWithdraw := &housetypes.MsgWithdraw{
		Creator:            owner.String(),
		MarketUID:          testMarketUID,
		Amount:             amt,
		ParticipationIndex: partecipationIndex,
		Mode:               housetypes.WithdrawalMode_WITHDRAWAL_MODE_FULL,
		Ticket:             ticket,
	}
	return inputWithdraw
}

func houseDepositMsg(t *testing.T, owner sdk.AccAddress, uid string, amt sdkmath.Int) *types.MsgHouseDeposit {
	testKyc := &sgetypes.KycDataPayload{
		Approved: true,
		ID:       owner.String(),
	}
	ticketClaim := jwt.MapClaims{
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
		"iat":      time.Now().Unix(),
		"kyc_data": testKyc,
	}
	ticket, err := simapp.CreateJwtTicket(ticketClaim)
	require.Nil(t, err)

	inputDeposit := &housetypes.MsgDeposit{
		Creator:   owner.String(),
		MarketUID: uid,
		Amount:    amt,
		Ticket:    ticket,
	}

	return &types.MsgHouseDeposit{
		Msg: inputDeposit,
	}
}

func testBetMsgWager(t testing.TB, bettor sdk.AccAddress, amount sdkmath.Int) *bettypes.MsgWager {
	betTicket, err := createJwtTicket(jwt.MapClaims{
		"exp":           9999999999,
		"iat":           7777777777,
		"selected_odds": testSelectedBetOdds,
		"kyc_data": &sgetypes.KycDataPayload{
			Approved: true,
			ID:       bettor.String(),
		},
		"all_odds": testBetOdds,
	})
	require.NoError(t, err)

	return &bettypes.MsgWager{
		Creator: bettor.String(),
		Props: &bettypes.WagerProps{
			UID:    uuid.NewString(),
			Amount: amount,
			Ticket: betTicket,
		},
	}
}
