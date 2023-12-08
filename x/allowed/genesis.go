package allowed

import (
	"github.com/andReyM228/one/x/allowed/keeper"
	"github.com/andReyM228/one/x/allowed/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the adresses
	for _, elem := range genState.AdressesList {
		k.SetAdresses(ctx, elem)
	}

	// Set adresses count
	k.SetAdressesCount(ctx, genState.AdressesCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AdressesList = k.GetAllAdresses(ctx)
	genesis.AdressesCount = k.GetAdressesCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
