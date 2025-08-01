package ctl

import (
	"context"
	"errors"
)

type Key int

var userKey Key

type UserInfo struct {
	Id uint `json:"id"`
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("获取用户信息失败")
	}
	return user, nil
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

func NewContext(ctx context.Context, user *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func InitUserInfo(ctx context.Context) {
	// TOOD 放缓存，之后的用户信息，走缓存
}
