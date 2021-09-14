package repository

import (
	"github.com/xiezerozero/mypackage/solid/dependency"
)

type MockUserinfo struct {
	Id           int
	AdEquityTime int64
}

func (m MockUserinfo) GetUserInfo(s string) (*dependency.UserInfoEntity, error) {
	return &dependency.UserInfoEntity{
		Id:           m.Id,
		AdEquityTime: m.AdEquityTime,
	}, nil
}
