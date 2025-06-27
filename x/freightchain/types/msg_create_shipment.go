package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = (*MsgCreateShipment)(nil)

func (msg *MsgCreateShipment) Route() string {
	return RouterKey
}

func (msg *MsgCreateShipment) Type() string {
	return "CreateShipment"
}

func (msg *MsgCreateShipment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateShipment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateShipment) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrAppConfig.Wrap(err.Error())
	}
	if msg.TrackingNumber == "" {
		return sdkerrors.ErrAppConfig.Wrap("tracking number cannot be empty")
	}
	if msg.Status == "" {
		msg.Status = string(StatusCreate)
	}
	return nil
}
