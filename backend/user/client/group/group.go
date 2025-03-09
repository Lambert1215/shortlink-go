// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package group

import (
	"context"

	"user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateGroupRequest     = user.CreateGroupRequest
	CreateGroupResponse    = user.CreateGroupResponse
	DeleteGroupRequest     = user.DeleteGroupRequest
	DeleteGroupResponse    = user.DeleteGroupResponse
	DeleteUserRequest      = user.DeleteUserRequest
	DeleteUserResponse     = user.DeleteUserResponse
	GetGroupRequest        = user.GetGroupRequest
	GetGroupResponse       = user.GetGroupResponse
	GetUserRequest         = user.GetUserRequest
	GetUserResponse        = user.GetUserResponse
	GroupData              = user.GroupData
	IsExistUserRequest     = user.IsExistUserRequest
	IsExistUserResponse    = user.IsExistUserResponse
	IsUserLoginRequest     = user.IsUserLoginRequest
	IsUserLoginResponse    = user.IsUserLoginResponse
	LoginRequest           = user.LoginRequest
	LoginResponse          = user.LoginResponse
	LogoutRequest          = user.LogoutRequest
	LogoutResponse         = user.LogoutResponse
	RegisterRequest        = user.RegisterRequest
	RegisterResponse       = user.RegisterResponse
	UpdateGroupRequest     = user.UpdateGroupRequest
	UpdateGroupResponse    = user.UpdateGroupResponse
	UpsertUserInfoRequest  = user.UpsertUserInfoRequest
	UpsertUserInfoResponse = user.UpsertUserInfoResponse

	Group interface {
		CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
		GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error)
		UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error)
		DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error)
	}

	defaultGroup struct {
		cli zrpc.Client
	}
)

func NewGroup(cli zrpc.Client) Group {
	return &defaultGroup{
		cli: cli,
	}
}

func (m *defaultGroup) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	client := user.NewGroupClient(m.cli.Conn())
	return client.CreateGroup(ctx, in, opts...)
}

func (m *defaultGroup) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*GetGroupResponse, error) {
	client := user.NewGroupClient(m.cli.Conn())
	return client.GetGroup(ctx, in, opts...)
}

func (m *defaultGroup) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*UpdateGroupResponse, error) {
	client := user.NewGroupClient(m.cli.Conn())
	return client.UpdateGroup(ctx, in, opts...)
}

func (m *defaultGroup) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*DeleteGroupResponse, error) {
	client := user.NewGroupClient(m.cli.Conn())
	return client.DeleteGroup(ctx, in, opts...)
}
