package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"

	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &productclient.RemoveRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.RemoveResponse{}, nil
}
