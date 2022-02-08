package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"

	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateLogic {
	return UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	if len(req.Name) == 0 || len(req.Desc) == 0 || req.Amount == 0 {
		return nil, errorx.NewDefaultError("params error.")
	}

	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &productclient.UpdateRequest{
		Id:     req.Id,
		Name:   req.Name,
		Desc:   req.Desc,
		Amount: req.Amount,
		Stock:  req.Stock,
		Status: req.Status,
	})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.UpdateResponse{}, nil
}
