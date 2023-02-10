package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/order/model"
	"github.com/honkkki/gomall/code/mall/service/product/rpc/types/product"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/types/user"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/order/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *order.CreateRequest) (*order.CreateResponse, error) {
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, err
	}

	// todo: 分布式事务
	productRes, err := l.svcCtx.ProductRpcClient.Detail(l.ctx, &product.DetailRequest{Id: in.Pid})
	if err != nil {
		return nil, err
	}
	if productRes.Stock <= 0 {
		return nil, status.Error(500, "stock not enough")
	}

	newOrder := model.Order{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: in.Amount,
		Status: 0,
	}
	// 创建订单
	res, err := l.svcCtx.OrderModel.Insert(l.ctx, &newOrder)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	newOrder.Id, err = res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新产品库存
	_, err = l.svcCtx.ProductRpcClient.UpdateStock(l.ctx, &product.UpdateStockRequest{
		Pid: in.Pid,
		Num: 1,
	})
	if err != nil {
		return nil, err
	}

	return &order.CreateResponse{
		Id: newOrder.Id,
	}, nil
}
