package logic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsExistUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsExistUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsExistUserLogic {
	return &IsExistUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// TODO 布隆过滤器
func (l *IsExistUserLogic) IsExistUser(in *user.IsExistUserRequest) (*user.IsExistUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.IsExistUserResponse{}, nil
}
