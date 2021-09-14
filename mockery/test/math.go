package test

type Totaller interface {
	Total(...int) int
}

type Calculator struct {
	total Totaller
}

func (c *Calculator) Double(inputs ...int) int {
	r := c.total.Total(inputs...)
	return r * 2
}
