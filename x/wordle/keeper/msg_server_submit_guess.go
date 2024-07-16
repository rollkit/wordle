package keeper

import (
	"context"

	"wordle/x/wordle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitGuess(goCtx context.Context, msg *types.MsgSubmitGuess) (*types.MsgSubmitGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitGuessResponse{}, nil
}
