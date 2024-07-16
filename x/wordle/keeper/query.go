package keeper

import (
	"wordle/x/wordle/types"
)

var _ types.QueryServer = Keeper{}
