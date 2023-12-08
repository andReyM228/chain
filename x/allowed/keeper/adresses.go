package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"one/x/allowed/types"
)

// GetAdressesCount get the total number of adresses
func (k Keeper) GetAdressesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdressesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAdressesCount set the total number of adresses
func (k Keeper) SetAdressesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AdressesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAdresses appends a adresses in the store with a new id and update the count
func (k Keeper) AppendAdresses(
	ctx sdk.Context,
	adresses types.Adresses,
) uint64 {
	// Create the adresses
	count := k.GetAdressesCount(ctx)

	// Set the ID of the appended value
	adresses.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	appendedValue := k.cdc.MustMarshal(&adresses)
	store.Set(GetAdressesIDBytes(adresses.Id), appendedValue)

	// Update adresses count
	k.SetAdressesCount(ctx, count+1)

	return count
}

// SetAdresses set a specific adresses in the store
func (k Keeper) SetAdresses(ctx sdk.Context, adresses types.Adresses) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	b := k.cdc.MustMarshal(&adresses)
	store.Set(GetAdressesIDBytes(adresses.Id), b)
}

// GetAdresses returns a adresses from its id
func (k Keeper) GetAdresses(ctx sdk.Context, id uint64) (val types.Adresses, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	b := store.Get(GetAdressesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAdressesByAdress returns a adresses from its adress
func (k Keeper) GetAdressesByAdress(ctx sdk.Context, address string) (val types.Adresses, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Adresses
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if val.Adress == address {
			return val, true
		}

	}
	return val, false
}

// RemoveAdresses removes a adresses from the store
func (k Keeper) RemoveAdresses(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	store.Delete(GetAdressesIDBytes(id))
}

// GetAllAdresses returns all adresses
func (k Keeper) GetAllAdresses(ctx sdk.Context) (list []types.Adresses) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AdressesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Adresses
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAdressesIDBytes returns the byte representation of the ID
func GetAdressesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAdressesIDFromBytes returns ID in uint64 format from a byte array
func GetAdressesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
