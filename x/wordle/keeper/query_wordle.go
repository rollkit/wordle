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

func (k Keeper) WordleAll(ctx context.Context, req *types.QueryAllWordleRequest) (*types.QueryAllWordleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var wordles []types.Wordle

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	wordleStore := prefix.NewStore(store, types.KeyPrefix(types.WordleKeyPrefix))

	pageRes, err := query.Paginate(wordleStore, req.Pagination, func(key []byte, value []byte) error {
		var wordle types.Wordle
		if err := k.cdc.Unmarshal(value, &wordle); err != nil {
			return err
		}

		wordles = append(wordles, wordle)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWordleResponse{Wordle: wordles, Pagination: pageRes}, nil
}

func (k Keeper) Wordle(ctx context.Context, req *types.QueryGetWordleRequest) (*types.QueryGetWordleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetWordle(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWordleResponse{Wordle: val}, nil
}
