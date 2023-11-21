package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"one/x/core/types"
)

// SetStats set a specific stats in the store from its index
func (k Keeper) SetStats(ctx sdk.Context, stats types.Stats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))
	b := k.cdc.MustMarshal(&stats)
	store.Set(types.StatsKey(
		stats.Index,
	), b)
}

// GetStats returns a stats from its index
func (k Keeper) GetStats(
	ctx sdk.Context,
	index string,

) (val types.Stats, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))

	b := store.Get(types.StatsKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStats removes a stats from the store
func (k Keeper) RemoveStats(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))
	store.Delete(types.StatsKey(
		index,
	))
}

// GetAllStats returns all stats
func (k Keeper) GetAllStats(ctx sdk.Context) (list []types.Stats) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stats
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
