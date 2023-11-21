package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "one/testutil/keeper"
	"one/testutil/nullify"
	"one/x/allowed/keeper"
	"one/x/allowed/types"
)

func createNAdresses(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Adresses {
	items := make([]types.Adresses, n)
	for i := range items {
		items[i].Id = keeper.AppendAdresses(ctx, items[i])
	}
	return items
}

func TestAdressesGet(t *testing.T) {
	keeper, ctx := keepertest.AllowedKeeper(t)
	items := createNAdresses(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAdresses(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAdressesRemove(t *testing.T) {
	keeper, ctx := keepertest.AllowedKeeper(t)
	items := createNAdresses(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdresses(ctx, item.Id)
		_, found := keeper.GetAdresses(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAdressesGetAll(t *testing.T) {
	keeper, ctx := keepertest.AllowedKeeper(t)
	items := createNAdresses(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdresses(ctx)),
	)
}

func TestAdressesCount(t *testing.T) {
	keeper, ctx := keepertest.AllowedKeeper(t)
	items := createNAdresses(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAdressesCount(ctx))
}
