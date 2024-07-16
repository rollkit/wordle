package keeper

import (
	"context"

	"wordle/x/wordle/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GuessAll(ctx context.Context, req *types.QueryAllGuessRequest) (*types.QueryAllGuessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var guesss []types.Guess

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	guessStore := prefix.NewStore(store, types.KeyPrefix(types.GuessKeyPrefix))

	pageRes, err := query.Paginate(guessStore, req.Pagination, func(key []byte, value []byte) error {
		var guess types.Guess
		if err := k.cdc.Unmarshal(value, &guess); err != nil {
			return err
		}

		guesss = append(guesss, guess)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGuessResponse{Guess: guesss, Pagination: pageRes}, nil
}

func (k Keeper) Guess(ctx context.Context, req *types.QueryGetGuessRequest) (*types.QueryGetGuessResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetGuess(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGuessResponse{Guess: val}, nil
}
