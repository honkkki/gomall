package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"

	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpcClient.Remove(l.ctx, &product.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.RemoveResponse{}, nil
}
