package repository

import (
	"time"
)

type MockCache struct {
	TimeUnix string
}

func (m MockCache) Get(key string) (string, error) {
	return m.TimeUnix, nil
}

func (m MockCache) Set(key string, value interface{}, expire time.Duration) {
}

func (m MockCache) Del(key string) error {
	return nil
}
