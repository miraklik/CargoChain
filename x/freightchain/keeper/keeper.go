package keeper

import (
	"fmt"

	"FreightChain/x/freightchain/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	corestore "cosmossdk.io/core/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeService corestore.KVStoreService
	cdc          codec.Codec
	addressCodec address.Codec
	// Address capable of executing a MsgUpdateParams message.
	// Typically, this should be the x/gov module account.
	authority []byte
	storeKey  storetypes.StoreKey

	Schema collections.Schema
	Params collections.Item[types.Params]
}

func NewKeeper(
	storeService corestore.KVStoreService,
	cdc codec.Codec,
	addressCodec address.Codec,
	authority []byte,

) Keeper {
	if _, err := addressCodec.BytesToString(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address %s: %s", authority, err))
	}

	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		storeService: storeService,
		cdc:          cdc,
		addressCodec: addressCodec,
		authority:    authority,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() []byte {
	return k.authority
}

func (k Keeper) SetShipment(ctx sdk.Context, shipment types.Shipment) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&shipment)
	store.Set(types.ShipmentKey(shipment.TrackingNumber), bz)
}

func (k Keeper) GetShipment(ctx sdk.Context, tracking string) (shipment types.Shipment, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ShipmentKey(tracking))
	if bz == nil {
		return shipment, false
	}
	k.cdc.MustUnmarshal(bz, &shipment)
	return shipment, true
}
