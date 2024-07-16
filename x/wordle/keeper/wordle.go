package keeper

import (
	"context"

	"wordle/x/wordle/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetWordle set a specific wordle in the store from its index
func (k Keeper) SetWordle(ctx context.Context, wordle types.Wordle) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WordleKeyPrefix))
	b := k.cdc.MustMarshal(&wordle)
	store.Set(types.WordleKey(
		wordle.Index,
	), b)
}

// GetWordle returns a wordle from its index
func (k Keeper) GetWordle(
	ctx context.Context,
	index string,

) (val types.Wordle, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WordleKeyPrefix))

	b := store.Get(types.WordleKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWordle removes a wordle from the store
func (k Keeper) RemoveWordle(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WordleKeyPrefix))
	store.Delete(types.WordleKey(
		index,
	))
}

// GetAllWordle returns all wordle
func (k Keeper) GetAllWordle(ctx context.Context) (list []types.Wordle) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WordleKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Wordle
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
