package tests

import (
	"strings"
	"testing"
)

var s = "310000,310000,510000,510100,510300,510400,510500,510600,510700,510800,510900,511000,511100"

func BenchmarkGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Split(s, ",")
	}
}
