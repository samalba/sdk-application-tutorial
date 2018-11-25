package poa

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewHandler returns a handler for "poa" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreateValidator:
			return handleMsgCreateValidator(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// EndBlocker is called for every block, update the validator set
func EndBlocker(ctx sdk.Context, k Keeper) []abci.ValidatorUpdate {
	updates := []abci.ValidatorUpdate{}
	// FIXME: All validators are returned from the store for each update.
	// Find a way to return only the ones that need to be updated (note: voting power == 0 means remove)
	validators := k.GetAllValidators(ctx)
	for _, val := range validators {
		updates = append(updates, val.ABCIValidatorUpdate())
	}
	return updates
}

// Handle MsgCreateValidator
func handleMsgCreateValidator(ctx sdk.Context, keeper Keeper, msg MsgCreateValidator) sdk.Result {
	if keeper.ValidatorExists(ctx, msg.ValidatorAddr) {
		return sdk.ErrInvalidAddress("Validator already exists with this address").Result()
	}
	keeper.CreateValidator(ctx, msg.ValidatorAddr, msg.PubKey)
	return sdk.Result{}
}
