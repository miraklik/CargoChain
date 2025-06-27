package keeper

import (
	"FreightChain/x/freightchain/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) CreateShipment(goCtx context.Context, msg *types.MsgCreateShipment) (*types.MsgCreateShipmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	shipment := types.NewShipment(msg.Creator, msg.TrackingNumber, msg.Status, msg.Location, ctx.BlockHeight())
	if _, found := m.GetShipment(ctx, msg.TrackingNumber); found {
		return nil, sdkerrors.ErrAppConfig.Wrap("shipment already exists")
	}

	m.SetShipment(ctx, *shipment)
	return &types.MsgCreateShipmentResponse{}, nil
}

func (m msgServer) UpdateShipment(goCtx context.Context, msg *types.MsgUpdateShipment) (*types.MsgUpdateShipmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	shipment, found := m.GetShipment(ctx, msg.TrackingNumber)
	if !found {
		return nil, sdkerrors.ErrAppConfig.Wrapf("shipment %s not found", msg.TrackingNumber)
	}

	if shipment.Creator != msg.Creator {
		return nil, sdkerrors.ErrAppConfig.Wrap("only creator can update their shipment")
	}

	shipment.Status = msg.Status
	shipment.Location = msg.Location
	shipment.Timestamp = ctx.BlockTime().Unix()

	m.SetShipment(ctx, shipment)
	return &types.MsgUpdateShipmentResponse{}, nil
}
