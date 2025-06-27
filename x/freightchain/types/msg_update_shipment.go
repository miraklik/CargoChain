package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = (*MsgUpdateShipment)(nil)

func (msg *MsgUpdateShipment) Route() string { return RouterKey }
func (msg *MsgUpdateShipment) Type() string  { return "UpdateShipment" }
func (msg *MsgUpdateShipment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
func (msg *MsgUpdateShipment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
func (msg *MsgUpdateShipment) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrAppConfig.Wrapf("invalid creator address (%s)", err)
	}
	if msg.TrackingNumber == "" {
		return sdkerrors.ErrAppConfig.Wrap("tracking number required")
	}
	return nil
}
