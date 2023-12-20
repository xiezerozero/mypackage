package finddifference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDifference(t *testing.T) {
	assert.Equal(t, FindDifference("abcd", "abcde"), "e")
	assert.Equal(t, FindDifference("", "t"), "t")
	assert.Equal(t, FindDifference("a", "aa"), "a")
	assert.Equal(t, FindDifference("ae", "aea"), "a")
}

func TestSingleNum(t *testing.T) {
	assert.Equal(t, SingleNum([]int{2, 2, 1}), 1)
	assert.Equal(t, SingleNum([]int{4, 1, 2, 1, 2}), 4)
}

func TestSortByBit(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}, SortByBit([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, SortByBit([]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}))
	assert.Equal(t, []int{10000, 10000}, SortByBit([]int{10000, 10000}))
	assert.Equal(t, []int{2, 3, 5, 17, 7, 11, 13, 19}, SortByBit([]int{2, 3, 5, 7, 11, 13, 17, 19}))
	assert.Equal(t, []int{10, 100, 10000, 1000}, SortByBit([]int{10, 100, 1000, 10000}))
}

func TestLowPower1(t *testing.T) {
	assert.Equal(t, 8, LowPower1(7))
	assert.Equal(t, 16, LowPower1(9))
	assert.Equal(t, 32, LowPower1(30))
	assert.Equal(t, 64, LowPower1(64))
}

func TestLowPower2(t *testing.T) {
	assert.Equal(t, 8, LowPower2(7))
	assert.Equal(t, 16, LowPower2(9))
	assert.Equal(t, 32, LowPower2(30))
	assert.Equal(t, 64, LowPower2(64))
}

func TestTwoSingleNums(t *testing.T) {
	assert.Equal(t, []int{3, 5}, TwoSingleNums([]int{1, 2, 1, 3, 2, 5}))
	assert.Equal(t, []int{1, 2}, TwoSingleNums([]int{1, 2, 3, 3, 4, 5, 6, 5, 4, 6}))
}

func TestOneCounts(t *testing.T) {
	assert.Equal(t, 2, OneCounts(9))
	assert.Equal(t, 3, OneCounts(35))
	assert.Equal(t, 2, OneCounts2(9))
	assert.Equal(t, 3, OneCounts2(35))
}

func TestHasAlternatingBits2(t *testing.T) {
	assert.Equal(t, false, HasAlternatingBits2(127)) // 1111111
	assert.Equal(t, false, HasAlternatingBits2(87))  // 1010111
	assert.Equal(t, true, HasAlternatingBits2(85))   // 1010101
}

func TestSwap(t *testing.T) {
	assert.Equal(t, []int{1, 2}, Swap([]int{2, 1}, 0, 1))
	assert.Equal(t, []int{1, 2}, Swap2([]int{2, 1}, 0, 1))
}
