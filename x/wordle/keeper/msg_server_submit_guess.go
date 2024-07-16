package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
	"wordle/x/wordle/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitGuess(goCtx context.Context, msg *types.MsgSubmitGuess) (*types.MsgSubmitGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check Word is 5 Characters Long
	if len(msg.Word) != 5 {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Guess Must Be A 5 Letter Word!")
	}

	// Check String Contains Alphabet Letters Only
	if !(IsLetter(msg.Word)) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Guess Must Only Consist of Alphabet Letters!")
	}

	// Get Current Day to Pull Up Wordle of That Day As A Hash
	currentTime := time.Now().Local()
	var currentTimeBytes = []byte(currentTime.Format("2006-01-02"))
	var currentTimeHash = sha256.Sum256(currentTimeBytes)
	var currentTimeHashString = hex.EncodeToString(currentTimeHash[:])
	wordle, isFound := k.GetWordle(ctx, currentTimeHashString)
	if !isFound {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "Wordle of The Day Hasn't Been Submitted Yet. Feel Free to Submit One!")
	}

	// We Convert Current Day and Guesser to A Hash To Use As An Index For Today's Guesses For That Guesser
	// That Way, A Person Can Guess 6 Times A Day For Each New Wordle Created
	var currentTimeGuesserBytes = []byte(currentTime.Format("2006-01-02") + msg.Creator)
	var currentTimeGuesserHash = sha256.Sum256(currentTimeGuesserBytes)
	var currentTimeGuesserHashString = hex.EncodeToString(currentTimeGuesserHash[:])
	// Hash The Guess To The Wordle
	var submittedSolutionHash = sha256.Sum256([]byte(msg.Word))
	var submittedSolutionHashString = hex.EncodeToString(submittedSolutionHash[:])

	// Get the Latest Guess entry for this Submitter for the current Wordle of the Day
	var count int
	guess, isFound := k.GetGuess(ctx, currentTimeGuesserHashString)
	if isFound {
		// Check if Submitter Reached 6 Tries
		if guess.Count == strconv.Itoa(6) {
			return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "You Have Guessed The Maximum Amount of Times for The Day! Try Again Tomorrow With A New Wordle.")
		}
		currentCount, err := strconv.Atoi(guess.Count)
		if err != nil {
			panic(err)
		}
		count = currentCount
	} else {
		// Initialize Count Value If No Entry Exists for this Submitter for Today's Wordle
		count = 0
	}
	// Increment Guess Count
	count += 1
	var newGuess = types.Guess{
		Index:     currentTimeGuesserHashString,
		Submitter: msg.Creator,
		Word:      submittedSolutionHashString,
		Count:     strconv.Itoa(count),
	}
	// Remove Current Guess Entry to be Updated With New Entry
	k.RemoveGuess(ctx, currentTimeGuesserHashString)
	// Add New Guess Entry
	k.SetGuess(ctx, newGuess)
	if !(wordle.Word == submittedSolutionHashString) {
		return &types.MsgSubmitGuessResponse{Title: "Wrong Answer", Body: "Your Guess Was Wrong. Try Again"}, nil
	} else {
		// Setup Reward
		reward := sdk.Coins{sdk.NewInt64Coin("token", 100)}
		// If Submitter Guesses Correctly
		guesserAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
		// Send Reward
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, guesserAddress, reward)
		if err != nil {
			return nil, err
		}
		return &types.MsgSubmitGuessResponse{Title: "Correct", Body: "You Guessed The Wordle Correctly!"}, nil
	}
}
