package di

type GoodsRepo interface {
	Get() Goods
}

type Goods struct {
}

func (g Goods) Name() string {
	return "name"
}
