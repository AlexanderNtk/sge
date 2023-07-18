package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sge-network/sge/utils"
)

const typeMsgResolveMarket = "market_resolve"

var _ sdk.Msg = &MsgResolve{}

// NewMsgResolve accepts the params to create new resolution body
func NewMsgResolve(creator, ticket string) *MsgResolve {
	return &MsgResolve{
		Creator: creator,
		Ticket:  ticket,
	}
}

// Route return the message route for slashing
func (*MsgResolve) Route() string { return RouterKey }

// Type return the resolve market type
func (*MsgResolve) Type() string { return typeMsgResolveMarket }

// GetSigners return the creators address
func (msg *MsgResolve) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes return the marshalled bytes of the msg
func (msg *MsgResolve) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validates the input resolution market
func (msg *MsgResolve) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Ticket == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ticket param")
	}

	return nil
}

// EmitEvent emits the event for the message success.
func (msg *MsgResolve) EmitEvent(ctx *sdk.Context, marketUID string) {
	emitter := utils.NewEventEmitter(ctx, attributeValueCategory)
	emitter.AddMsg(typeMsgResolveMarket, msg.Creator,
		sdk.NewAttribute(attributeKeyMarketUID, marketUID),
	)
	emitter.Emit()
}
