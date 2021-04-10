package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/swap/types"
)

var (
	_ types.MsgServiceServer = server{}
)

type server struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &server{Keeper: keeper}
}

func (k server) MsgSwap(c context.Context, msg *types.MsgSwapRequest) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	if !k.SwapEnabled(ctx) {
		return nil, types.ErrorSwapIsDisabled
	}
	if k.ApproveBy(ctx) != msg.From {
		return nil, types.ErrorUnauthorized
	}
	if k.HasSwap(ctx, types.BytesToHash(msg.TxHash)) {
		return nil, types.ErrorDuplicateSwap
	}

	var (
		coin = sdk.NewCoin(k.SwapDenom(ctx), msg.Amount.Quo(types.PrecisionLoss))
		swap = types.Swap{
			TxHash:   msg.TxHash,
			Receiver: msg.Receiver,
			Amount:   coin,
		}
	)

	msgReceiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}

	if err := k.MintCoin(ctx, coin); err != nil {
		return nil, err
	}
	if err := k.SendCoinFromModuleToAccount(ctx, msgReceiver, coin); err != nil {
		return nil, err
	}

	k.SetSwap(ctx, swap)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSet,
		sdk.NewAttribute(types.AttributeKeyTxHash, fmt.Sprintf("%X", swap.TxHash)),
		sdk.NewAttribute(types.AttributeKeyAddress, swap.Receiver),
		sdk.NewAttribute(types.AttributeKeyAmount, swap.Amount.String()),
	))

	return &types.MsgSwapResponse{}, nil
}
