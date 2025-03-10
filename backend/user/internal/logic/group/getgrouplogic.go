package grouplogic

import (
	"context"

	"user/internal/svc"
	"user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupLogic {
	return &GetGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupLogic) GetGroup(in *user.GetGroupRequest) (*user.GetGroupResponse, error) {
	groupInfo, err := l.svcCtx.GroupModel.GetGroupByGid(l.ctx, in.GetGid())
	if err != nil {
		return nil, err
	}

	return &user.GetGroupResponse{
		Gid:      groupInfo.GID,
		Name:     groupInfo.Name,
		Username: groupInfo.Username,
	}, nil
}
