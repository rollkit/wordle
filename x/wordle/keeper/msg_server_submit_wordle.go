package keeper

import (
	"context"

	"wordle/x/wordle/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitWordle(goCtx context.Context, msg *types.MsgSubmitWordle) (*types.MsgSubmitWordleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSubmitWordleResponse{}, nil
}
