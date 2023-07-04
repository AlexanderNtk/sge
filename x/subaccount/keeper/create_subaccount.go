package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sge-network/sge/app/params"
	"github.com/sge-network/sge/x/subaccount/types"
)

func (m msgServer) CreateSubAccount(
	ctx context.Context,
	request *types.MsgCreateSubAccountRequest,
) (*types.MsgCreateAccountResponse, error) {
	sdkContext := sdk.UnwrapSDKContext(ctx)
	moneyToSend := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, sdk.NewInt(0)))

	err := request.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "invalid request")
	}

	for _, balanceUnlock := range request.LockedBalances {
		if balanceUnlock.UnlockTime.Unix() < sdkContext.BlockTime().Unix() {
			return nil, types.ErrUnlockTokenTimeExpired
		}

		moneyToSend = moneyToSend.Add(sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, balanceUnlock.Amount))...)
	}

	senderAccount, _ := sdk.AccAddressFromBech32(request.Sender)
	subaccountOwner, _ := sdk.AccAddressFromBech32(request.SubAccountOwner)
	if m.keeper.HasSubAccount(sdkContext, subaccountOwner) {
		return nil, types.ErrSubaccountAlreadyExist
	}

	m.keeper.NextID(sdkContext)
	subaccountID := m.keeper.Peek(sdkContext)

	subaccountAddress := types.NewAddressFromSubaccount(subaccountID)
	address := m.accountKeeper.NewAccountWithAddress(sdkContext, subaccountAddress)
	m.accountKeeper.SetAccount(sdkContext, address)

	err = m.bankKeeper.SendCoins(sdkContext, senderAccount, subaccountAddress, moneyToSend)
	if err != nil {
		return nil, errors.Wrap(err, "unable to send coins")
	}

	m.keeper.SetSubAccountOwner(sdkContext, subaccountID, subaccountOwner)
	m.keeper.SetLockedBalances(sdkContext, subaccountAddress, request.LockedBalances)

	return &types.MsgCreateAccountResponse{}, nil
}
