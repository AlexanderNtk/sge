package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/sge-network/sge/testutil/nullify"
	"github.com/sge-network/sge/testutil/sample"
	"github.com/stretchr/testify/require"

	"github.com/sge-network/sge/x/house/keeper"
	"github.com/sge-network/sge/x/house/types"
)

func createNDeposits(
	keeper *keeper.KeeperTest,
	ctx sdk.Context,
	n int,
) []types.Deposit {
	items := make([]types.Deposit, n)

	for i := range items {
		items[i].Creator = sample.AccAddress()
		items[i].MarketUID = testMarketUID
		items[i].ParticipationIndex = uint64(i + 1)
		items[i].DepositorAddress = testDepositorAddress
		items[i].Amount = sdkmath.NewInt(100)
		items[i].TotalWithdrawalAmount = sdkmath.NewInt(0)

		keeper.SetDeposit(ctx, items[i])
	}
	return items
}

func TestDepositGet(t *testing.T) {
	k, ctx := setupKeeper(t)
	items := createNDeposits(k, ctx, 10)

	rst, found := k.GetDeposit(ctx,
		uuid.NewString(),
		uuid.NewString(),
		1,
	)
	var expectedResp types.Deposit
	require.False(t, found)
	require.Equal(t,
		nullify.Fill(expectedResp),
		nullify.Fill(rst),
	)

	for _, item := range items {
		rst, found := k.GetDeposit(ctx,
			item.DepositorAddress,
			item.MarketUID,
			item.ParticipationIndex,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(item),
			nullify.Fill(rst),
		)
	}
}

func TestDepositGetAll(t *testing.T) {
	k, ctx := setupKeeper(t)
	items := createNDeposits(k, ctx, 10)

	deposits, err := k.GetAllDeposits(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(deposits),
	)
}
