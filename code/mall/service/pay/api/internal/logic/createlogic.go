package logic

import (
	"context"

	"github.com/honkkki/gomall/code/mall/service/pay/rpc/types/pay"

	"github.com/honkkki/gomall/code/mall/service/pay/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/pay/api/internal/types"

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
	res, err := l.svcCtx.PayRpcClient.Create(l.ctx, &pay.CreateRequest{
		Uid:    req.Uid,
		Oid:    req.Oid,
		Amount: req.Amount,
	})

	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	return &types.CreateResponse{
		Id: res.Id,
	}, nil
}
