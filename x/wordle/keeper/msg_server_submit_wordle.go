package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"unicode"
	"wordle/x/wordle/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitWordle(goCtx context.Context, msg *types.MsgSubmitWordle) (*types.MsgSubmitWordleResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check to See the Wordle is 5 letters
	if len(msg.Word) != 5 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Wordle Must Be A 5 Letter Word")
	}
	// Check to See Only Alphabets Are Passed for the Wordle
	if !(IsLetter(msg.Word)) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Wordle Must Only Consist Of Letters In The Alphabet")
	}

	// Use Current Day to Create The Index of the Newly-Submitted Wordle of the Day
	currentTime := time.Now().Local()
	var currentTimeBytes = []byte(currentTime.Format("2006-01-02"))
	var currentTimeHash = sha256.Sum256(currentTimeBytes)
	var currentTimeHashString = hex.EncodeToString(currentTimeHash[:])
	// Hash The Newly-Submitted Wordle of the Day
	var submittedSolutionHash = sha256.Sum256([]byte(msg.Word))
	var submittedSolutionHashString = hex.EncodeToString(submittedSolutionHash[:])

	var wordle = types.Wordle{
		Index:     currentTimeHashString,
		Word:      submittedSolutionHashString,
		Submitter: msg.Creator,
	}

	// Try to Get Wordle From KV Store Using Current Day as Key
	// This Helps ensure only one Wordle is submitted per day
	_, isFound := k.GetWordle(ctx, currentTimeHashString)
	if isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Wordle of the Day is Already Submitted")
	}
	// Write Wordle to KV Store
	k.SetWordle(ctx, wordle)
	reward := sdk.Coins{sdk.NewInt64Coin("token", 100)}
	// Escrow Reward
	submitterAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, submitterAddress, types.ModuleName, reward)
	if err != nil {
		return nil, err
	}
	return &types.MsgSubmitWordleResponse{}, nil
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
