package repository

import (
	"github.com/xiezerozero/mypackage/solid/dependency"
)

type RedisUserInfo struct {
}

func (r RedisUserInfo) GetUserInfo(s string) (*dependency.UserInfoEntity, error) {
	panic("implement me")
}
