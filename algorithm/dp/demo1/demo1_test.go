package demo1

import (
	"fmt"
	"testing"
)

func TestMaxValue(t *testing.T) {
	d1 := [][]int{
		{1, 8, 3, 4},
		{5, 6, 7, 9},
	}

	fmt.Println(MaxValue(d1))
	fmt.Println(MaxValue2(d1))
}
