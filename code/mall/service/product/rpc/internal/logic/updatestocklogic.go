package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/product/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStockLogic {
	return &UpdateStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStockLogic) UpdateStock(in *product.UpdateStockRequest) (*product.UpdateStockResponse, error) {
	err := l.svcCtx.ProductModel.UpdateStock(l.ctx, in.Pid, in.Num)
	if err != nil {
		return nil, status.Error(500, fmt.Sprintf("update stock failed: %s", err.Error()))
	}

	return &product.UpdateStockResponse{}, nil
}
