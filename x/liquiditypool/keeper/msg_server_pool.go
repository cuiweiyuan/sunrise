package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pool = types.Pool{
		Admin:      msg.Creator,
		BaseDenom:  msg.BaseDenom,
		QuoteDenom: msg.QuoteDenom,
	}

	id := k.AppendPool(
		ctx,
		pool,
	)

	k.CheckSetInitialPairAndTwap(ctx, pool.BaseDenom, pool.QuoteDenom)

	return &types.MsgCreatePoolResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePool(goCtx context.Context, msg *types.MsgUpdatePool) (*types.MsgUpdatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var pool = types.Pool{
		Admin:      msg.Admin,
		Id:         msg.Id,
		BaseDenom:  msg.BaseDenom,
		QuoteDenom: msg.QuoteDenom,
	}

	// Checks that the element exists
	val, found := k.GetPool(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current admin
	if msg.Admin != val.Admin {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect admin")
	}

	if msg.NewAdmin != "" {
		pool.Admin = msg.NewAdmin
	}

	k.SetPool(ctx, pool)

	return &types.MsgUpdatePoolResponse{}, nil
}
