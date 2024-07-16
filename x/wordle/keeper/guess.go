package keeper

import (
	"context"

	"wordle/x/wordle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetGuess set a specific guess in the store from its index
func (k Keeper) SetGuess(ctx context.Context, guess types.Guess) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKeyPrefix))
	b := k.cdc.MustMarshal(&guess)
	store.Set(types.GuessKey(
		guess.Index,
	), b)
}

// GetGuess returns a guess from its index
func (k Keeper) GetGuess(
	ctx context.Context,
	index string,

) (val types.Guess, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKeyPrefix))

	b := store.Get(types.GuessKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveGuess removes a guess from the store
func (k Keeper) RemoveGuess(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKeyPrefix))
	store.Delete(types.GuessKey(
		index,
	))
}

// GetAllGuess returns all guess
func (k Keeper) GetAllGuess(ctx context.Context) (list []types.Guess) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.GuessKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Guess
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
