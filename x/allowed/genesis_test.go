package allowed_test

import (
	"testing"

	keepertest "github.com/andReyM228/one/testutil/keeper"
	"github.com/andReyM228/one/testutil/nullify"
	"github.com/andReyM228/one/x/allowed"
	"github.com/andReyM228/one/x/allowed/types"
	"github.com/stretchr/testify/require"
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
