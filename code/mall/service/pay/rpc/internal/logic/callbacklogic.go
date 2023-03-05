package logic

import (
	"context"

	"github.com/honkkki/gomall/code/mall/service/order/rpc/types/order"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/types/user"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/pay/model"
	"github.com/honkkki/gomall/code/mall/service/pay/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/pay/rpc/types/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpcClient.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	detail, err := l.svcCtx.OrderRpcClient.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, err
	}

	if detail.Uid != in.Uid {
		return nil, status.Error(100, "订单不属于该用户")
	}

	// 查询支付是否存在
	res, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 支付金额与订单金额不符
	if in.Amount != res.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	res.Source = in.Source
	res.Status = in.Status

	err = l.svcCtx.PayModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpcClient.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &pay.CallbackResponse{}, nil
}
