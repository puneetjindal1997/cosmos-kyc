package keeper

import (
	"context"

	"kyc/x/kyc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateKyc(goCtx context.Context, msg *types.MsgCreateKyc) (*types.MsgCreateKycResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message

	kyc := types.Kyc{
		Creator:     msg.Creator,
		Doc:         msg.Doc,
		Description: msg.Description,
	}
	id := k.AppendKyc(ctx, kyc)

	return &types.MsgCreateKycResponse{Id: id}, nil
}
