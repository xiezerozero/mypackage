package contains

import "testing"

func Benchmark_contains1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 调用你要测试的函数
		result := contains1(Ids, Input)
		_ = result
	}
}

func Benchmark_contains2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 调用你要测试的函数
		result := contains2(Ids, Input)
		_ = result
	}
}
