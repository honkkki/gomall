package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/service/user/model"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/user/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	res, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "user not found.")
		}
		return nil, status.Error(500, err.Error())
	}

	return &user.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
