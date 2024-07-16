package wordle

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"wordle/testutil/sample"
	wordlesimulation "wordle/x/wordle/simulation"
	"wordle/x/wordle/types"
)

// avoid unused import issue
var (
	_ = wordlesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSubmitWordle = "op_weight_msg_submit_wordle"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitWordle int = 100

	opWeightMsgSubmitGuess = "op_weight_msg_submit_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGuess int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	wordleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&wordleGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSubmitWordle int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitWordle, &weightMsgSubmitWordle, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitWordle = defaultWeightMsgSubmitWordle
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitWordle,
		wordlesimulation.SimulateMsgSubmitWordle(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitGuess int
	simState.AppParams.GetOrGenerate(opWeightMsgSubmitGuess, &weightMsgSubmitGuess, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitGuess = defaultWeightMsgSubmitGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitGuess,
		wordlesimulation.SimulateMsgSubmitGuess(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitWordle,
			defaultWeightMsgSubmitWordle,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				wordlesimulation.SimulateMsgSubmitWordle(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSubmitGuess,
			defaultWeightMsgSubmitGuess,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				wordlesimulation.SimulateMsgSubmitGuess(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
