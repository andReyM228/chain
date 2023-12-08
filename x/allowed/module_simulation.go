package allowed

import (
	"math/rand"

	"github.com/andReyM228/one/testutil/sample"
	allowedsimulation "github.com/andReyM228/one/x/allowed/simulation"
	"github.com/andReyM228/one/x/allowed/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = allowedsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateAdresses = "op_weight_msg_adresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAdresses int = 100

	opWeightMsgUpdateAdresses = "op_weight_msg_adresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAdresses int = 100

	opWeightMsgDeleteAdresses = "op_weight_msg_adresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAdresses int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	allowedGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AdressesList: []types.Adresses{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AdressesCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&allowedGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAdresses int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAdresses, &weightMsgCreateAdresses, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAdresses = defaultWeightMsgCreateAdresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAdresses,
		allowedsimulation.SimulateMsgCreateAdresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAdresses int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAdresses, &weightMsgUpdateAdresses, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAdresses = defaultWeightMsgUpdateAdresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAdresses,
		allowedsimulation.SimulateMsgUpdateAdresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAdresses int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAdresses, &weightMsgDeleteAdresses, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAdresses = defaultWeightMsgDeleteAdresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAdresses,
		allowedsimulation.SimulateMsgDeleteAdresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAdresses,
			defaultWeightMsgCreateAdresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				allowedsimulation.SimulateMsgCreateAdresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAdresses,
			defaultWeightMsgUpdateAdresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				allowedsimulation.SimulateMsgUpdateAdresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAdresses,
			defaultWeightMsgDeleteAdresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				allowedsimulation.SimulateMsgDeleteAdresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
