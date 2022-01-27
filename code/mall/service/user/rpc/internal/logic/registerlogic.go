package logic

import (
	"context"
	"github.com/honkkki/gomall/code/mall/common/encryption"
	"github.com/honkkki/gomall/code/mall/service/user/model"
	"google.golang.org/grpc/status"

	"github.com/honkkki/gomall/code/mall/service/user/rpc/internal/svc"
	"github.com/honkkki/gomall/code/mall/service/user/rpc/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 用户注册的具体逻辑
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err == nil {
		return nil, status.Error(100, "user mobile has been used.")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   in.Gender,
			Mobile:   in.Mobile,
			Password: encryption.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(&newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		id, err := res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.RegisterResponse{
			Id:     id,
			Name:   newUser.Name,
			Gender: newUser.Gender,
			Mobile: newUser.Mobile,
		}, nil
	}

	return nil, status.Error(500, err.Error())
}
