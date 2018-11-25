package poa

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// QueryValidators is used for the API path
	QueryValidators = "validators"
)

// NewQuerier creates a querier for the REST endpoint
func NewQuerier(k Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryValidators:
			return queryValidators(ctx, cdc, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown stake query endpoint")
		}
	}
}

func queryValidators(ctx sdk.Context, cdc *codec.Codec, k Keeper) (res []byte, err sdk.Error) {
	validators := k.GetAllValidators(ctx)
	res, errRes := codec.MarshalJSONIndent(cdc, validators)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", errRes.Error()))
	}
	return res, nil
}
