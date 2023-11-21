package one_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "one/testutil/keeper"
	"one/testutil/nullify"
	"one/x/one"
	"one/x/one/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OneKeeper(t)
	one.InitGenesis(ctx, *k, genesisState)
	got := one.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
