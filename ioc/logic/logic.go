package logic

import "github.com/xiezerozero/mypackage/ioc/di"

type Logic struct {
	gp di.GoodsRepo
}

func (l Logic) Get() {
	g := l.gp.Get()
	g.Name()

}
