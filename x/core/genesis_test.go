package core_test

import (
	"testing"

	keepertest "github.com/andReyM228/one/testutil/keeper"
	"github.com/andReyM228/one/testutil/nullify"
	"github.com/andReyM228/one/x/core"
	"github.com/andReyM228/one/x/core/types"
	"github.com/stretchr/testify/require"
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
