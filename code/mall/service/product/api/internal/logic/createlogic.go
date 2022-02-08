package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/api/internal/types"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/productclient"
	"github.com/tal-tech/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type CreateRequest struct {
	Name   string  `json:"name"`
	Desc   string  `json:"desc"`
	Stock  int64   `json:"stock"`
	Amount float64 `json:"amount"`
	Status int64   `json:"status"`
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateLogic {
	return CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req types.CreateRequest) (resp *types.CreateResponse, err error) {
	if len(req.Name) == 0 || len(req.Desc) == 0 || req.Amount == 0 {
		return nil, errorx.NewDefaultError("params error.")
	}

	res, err := l.svcCtx.ProductRpc.Create(l.ctx, &productclient.CreateRequest{
		Name:   req.Name,
		Desc:   req.Desc,
		Stock:  req.Stock,
		Amount: req.Amount,
		Status: req.Status,
	})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.CreateResponse{
		Id:   res.Id,
		Name: res.Name,
	}, nil
}
