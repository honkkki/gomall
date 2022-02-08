package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/product/model"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/product"
	"google.golang.org/grpc/status"

	"github.com/tal-tech/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	var num int64
	err := l.svcCtx.DBEngine.Model(&model.Product{}).Where("id=?", in.Id).Count(&num).Error
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	if num == 0 {
		return nil, status.Error(100, "product not exist.")
	}

	err = l.svcCtx.ProductModel.Update(&model.Product{
		Id:     in.Id,
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.UpdateResponse{}, nil
}
