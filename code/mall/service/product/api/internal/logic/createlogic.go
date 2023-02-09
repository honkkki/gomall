package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"

	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	res, err := l.svcCtx.ProductRpcClient.Create(l.ctx, &product.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateResponse{
		Id:   res.Id,
		Name: res.Name,
	}, nil
}
