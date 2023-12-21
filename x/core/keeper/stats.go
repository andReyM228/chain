package keeper

import (
	"github.com/andReyM228/one/x/core/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"
)

func (k Keeper) SaveStatsIssue(ctx sdk.Context, amount sdk.Coins) {
	var (
		store       = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))
		amountIssue sdk.Coins
		currentDate = time.Now().Format(time.DateOnly)
	)

	stat, found := k.GetStatsByDate(ctx, currentDate)
	if !found {
		stats := k.GetAllStats(ctx)

		indexStats := len(stats) + 1

		stat = types.Stats{
			Index: strconv.Itoa(indexStats),
			Date:  currentDate,
			Stats: &types.DailyStats{
				AmountIssue:    sdk.Coins{},
				AmountWithdraw: sdk.Coins{},
				CountIssue:     0,
				CountWithdraw:  0,
			},
		}
	}

	amountIssue = stat.Stats.AmountIssue

	stat.Stats.CountIssue += 1
	stat.Stats.AmountIssue = amountIssue.Add(amount...)

	result := k.cdc.MustMarshal(&stat)

	store.Set(types.StatsKey(stat.Index), result)
}

func (k Keeper) SaveStatsWithdraw(ctx sdk.Context, amount sdk.Coins) {
	var (
		store          = prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))
		amountWithdraw sdk.Coins
		currentDate    = time.Now().Format(time.DateOnly)
	)

	stat, found := k.GetStatsByDate(ctx, currentDate)
	if !found {
		stats := k.GetAllStats(ctx)

		indexStats := len(stats) + 1

		stat = types.Stats{
			Index: strconv.Itoa(indexStats),
			Date:  currentDate,
			Stats: &types.DailyStats{
				AmountIssue:    sdk.Coins{},
				AmountWithdraw: sdk.Coins{},
				CountIssue:     0,
				CountWithdraw:  0,
			},
		}
	}

	amountWithdraw = stat.Stats.AmountWithdraw

	stat.Stats.CountWithdraw += 1
	stat.Stats.AmountWithdraw = amountWithdraw.Add(amount...)

	result := k.cdc.MustMarshal(&stat)

	store.Set(types.StatsKey(stat.Index), result)
}

func (k Keeper) GetStatsByDate(ctx sdk.Context, date string) (stats types.Stats, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatsKeyPrefix))

	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	parsedDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return stats, false
	}

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stats
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		statsDate, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			continue
		}

		if statsDate.Equal(parsedDate) {
			return val, true
		}
	}

	return stats, false
}

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
