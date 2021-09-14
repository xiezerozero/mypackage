package db

import "github.com/xiezerozero/mypackage/ioc/di"

type Db struct {
}

func (db Db) Get() di.Goods {
	return di.Goods{}
}
