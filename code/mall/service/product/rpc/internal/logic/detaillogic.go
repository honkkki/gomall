package logic

import (
	"context"
	"errors"
	"github.com/honkkki/gomall/code/mall/service/product/model"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
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
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "product not found")
		}
		return nil, status.Error(500, err.Error())
	}
	l.svcCtx.RedisClient.Set("name", "zero")
	logx.Info("it is detail rpc method")

	return &product.DetailResponse{
		Id:     res.Id,
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: res.Amount,
		Status: res.Status,
	}, nil
}
