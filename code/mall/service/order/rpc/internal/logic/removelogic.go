package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/order/model"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/order/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *order.RemoveRequest) (*order.RemoveResponse, error) {
	// 查询订单是否存在
	res, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.OrderModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.RemoveResponse{}, nil
}
