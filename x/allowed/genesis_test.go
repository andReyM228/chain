package allowed_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "one/testutil/keeper"
	"one/testutil/nullify"
	"one/x/allowed"
	"one/x/allowed/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AdressesList: []types.Adresses{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AdressesCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AllowedKeeper(t)
	allowed.InitGenesis(ctx, *k, genesisState)
	got := allowed.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AdressesList, got.AdressesList)
	require.Equal(t, genesisState.AdressesCount, got.AdressesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
