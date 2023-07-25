package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	bettypes "github.com/sge-network/sge/x/bet/types"
	housetypes "github.com/sge-network/sge/x/house/types"
	orderbookmodulekeeper "github.com/sge-network/sge/x/orderbook/keeper"
	"github.com/sge-network/sge/x/subaccount/types"
)

type BetKeeper interface {
	GetBetID(ctx sdk.Context, uid string) (bettypes.UID2ID, bool)
	PlaceBet(ctx sdk.Context, bet *bettypes.Bet) error
}

type BankKeeper interface {
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

type HouseKeeper interface {
	GetParams(ctx sdk.Context) housetypes.Params
	Deposit(ctx sdk.Context, creator, depositor, marketUID string, amount sdk.Int) (participationIndex uint64, err error)
}

type OrderBookKeeper interface {
	RegisterHook(hooks orderbookmodulekeeper.Hook)
}

type Keeper struct {
	cdc        codec.Codec
	storeKey   sdk.StoreKey
	paramstore paramtypes.Subspace
	bankKeeper BankKeeper

	ovmKeeper   bettypes.OVMKeeper
	betKeeper   BetKeeper
	houseKeeper HouseKeeper
}

func NewKeeper(cdc codec.Codec, storeKey sdk.StoreKey, ps paramtypes.Subspace, bankKeeper BankKeeper, ovmKeeper bettypes.OVMKeeper, betKeeper BetKeeper, obKeeper OrderBookKeeper, hk HouseKeeper) Keeper {
	// set KeyTable if it is not already set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	k := Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		paramstore:  ps,
		bankKeeper:  bankKeeper,
		ovmKeeper:   ovmKeeper,
		betKeeper:   betKeeper,
		houseKeeper: hk,
	}
	obKeeper.RegisterHook(k)
	return k
}
