package core_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "one/testutil/keeper"
	"one/testutil/nullify"
	"one/x/core"
	"one/x/core/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StatsList: []types.Stats{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	core.InitGenesis(ctx, *k, genesisState)
	got := core.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StatsList, got.StatsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
