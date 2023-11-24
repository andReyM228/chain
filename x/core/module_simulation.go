package core

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"one/testutil/sample"
	coresimulation "one/x/core/simulation"
	"one/x/core/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = coresimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStats int = 100

	opWeightMsgUpdateStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStats int = 100

	opWeightMsgDeleteStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStats int = 100

	opWeightMsgIssue = "op_weight_msg_issue"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssue int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	coreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StatsList: []types.Stats{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&coreGenesis)
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

	var weightMsgCreateStats int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateStats, &weightMsgCreateStats, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStats = defaultWeightMsgCreateStats
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateStats,
		coresimulation.SimulateMsgCreateStats(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateStats int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateStats, &weightMsgUpdateStats, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStats = defaultWeightMsgUpdateStats
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateStats,
		coresimulation.SimulateMsgUpdateStats(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteStats int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteStats, &weightMsgDeleteStats, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStats = defaultWeightMsgDeleteStats
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteStats,
		coresimulation.SimulateMsgDeleteStats(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIssue int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgIssue, &weightMsgIssue, nil,
		func(_ *rand.Rand) {
			weightMsgIssue = defaultWeightMsgIssue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssue,
		coresimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		coresimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateStats,
			defaultWeightMsgCreateStats,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgCreateStats(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateStats,
			defaultWeightMsgUpdateStats,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgUpdateStats(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteStats,
			defaultWeightMsgDeleteStats,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgDeleteStats(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgIssue,
			defaultWeightMsgIssue,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdraw,
			defaultWeightMsgWithdraw,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
