package model

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user/pkg/errorcode"
)

const (
	DB         = "shortlink"
	Collection = "user"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Register(ctx context.Context, data *User) error
		UpdateUserInfo(ctx context.Context, data *User) error
		GetUserByUserName(ctx context.Context, username string) (*User, error)
		Login(ctx context.Context, username, password string) bool
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (c *customUserModel) Login(ctx context.Context, username, password string) bool {
	filter := bson.M{
		"username":   username,
		"password":   password,
		"deleteFlag": 0,
	}

	var user User
	err := c.conn.FindOne(ctx, &user, filter)
	if err != nil {
		return false
	}
	return true
}

func (c *customUserModel) GetUserByUserName(ctx context.Context, username string) (*User, error) {
	filter := bson.M{"username": username}

	var user User

	err := c.conn.FindOne(ctx, user, filter)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, errors.New(errorcode.UserNotExist.Message())
		}
		return nil, err
	}

	return &user, nil
}

func (c *customUserModel) Register(ctx context.Context, data *User) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	filter := bson.M{
		"username":   data.Username,
		"deleteFlag": bson.M{"$ne": 1},
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New(errorcode.UserNameExistError.Message())
	}

	_, err = c.conn.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *customUserModel) UpdateUserInfo(ctx context.Context, data *User) error {
	data.UpdateAt = time.Now()

	filter := bson.D{
		{"username", data.Username},
	}
	update := bson.D{
		{"$set", data},
	}
	_, err := c.conn.UpdateOne(ctx, filter, update)
	return err
}

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string) UserModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}
