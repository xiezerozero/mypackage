package main

import (
	"fmt"
	"github.com/xiezerozero/go-generic/slices"
)

func main() {
	result := slices.Filter([]int{1, 10, 100, 1000, 2, 5, 102}, func(v int) bool {
		return v > 100
	})
	fmt.Println(result)

	stringResult := slices.Filter([]string{"a", "b", "c", "d", "e"}, func(v string) bool {
		return v > "c"
	})

	fmt.Println(stringResult)
}
