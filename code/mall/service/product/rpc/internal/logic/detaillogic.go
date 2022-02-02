package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/user/model"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/product"

	"github.com/tal-tech/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	res, err := l.svcCtx.ProductModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "product not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &product.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
