package wordle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"wordle/x/wordle/keeper"
	"wordle/x/wordle/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the wordle
	for _, elem := range genState.WordleList {
		k.SetWordle(ctx, elem)
	}
	// Set all the guess
	for _, elem := range genState.GuessList {
		k.SetGuess(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.WordleList = k.GetAllWordle(ctx)
	genesis.GuessList = k.GetAllGuess(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
