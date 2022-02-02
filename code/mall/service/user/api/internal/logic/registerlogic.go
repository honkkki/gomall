package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/errorx"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/userclient"
	"strings"

	"github.com/honkkki/gomall/code/mall/service/user/api/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	if len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.Password)) == 0 || len(strings.TrimSpace(req.Mobile)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	res, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})

	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.RegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
