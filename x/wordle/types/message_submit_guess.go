package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitGuess{}

func NewMsgSubmitGuess(creator string, word string) *MsgSubmitGuess {
	return &MsgSubmitGuess{
		Creator: creator,
		Word:    word,
	}
}

func (msg *MsgSubmitGuess) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
