package logic

import (
	"context"
	model "user/mongo/user"
	"user/pkg/constant"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	err := l.svcCtx.UserModel.Register(l.ctx, &model.User{
		Username:   in.Username,
		Password:   in.Password,
		RealName:   in.RealName,
		Phone:      in.Phone,
		Mail:       in.Email,
		DeleteFlag: constant.DELETE_FLAG,
	})
	if err != nil {
		return &user.RegisterResponse{
			Success: false,
		}, err
	}
	return &user.RegisterResponse{
		Success: true,
	}, nil
}
