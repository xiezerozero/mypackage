package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	div3()
}

func div3() {
	d := decimal.NewFromInt(1)
	divD := d.DivRound(decimal.NewFromInt(3), 2)
	fmt.Println(divD)
	fmt.Println(d.Sub(divD).Sub(divD))
}

func round() {
	d, err := decimal.NewFromString("100.1216")
	fmt.Println(d, err)

	d = d.Mul(decimal.NewFromInt(100))
	fmt.Println(d.Round(0).IntPart())
}
