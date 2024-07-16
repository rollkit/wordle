package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "wordle/testutil/keeper"
	"wordle/testutil/nullify"
	"wordle/x/wordle/keeper"
	"wordle/x/wordle/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNGuess(keeper keeper.Keeper, ctx context.Context, n int) []types.Guess {
	items := make([]types.Guess, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetGuess(ctx, items[i])
	}
	return items
}

func TestGuessGet(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetGuess(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestGuessRemove(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveGuess(ctx,
			item.Index,
		)
		_, found := keeper.GetGuess(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestGuessGetAll(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNGuess(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllGuess(ctx)),
	)
}
