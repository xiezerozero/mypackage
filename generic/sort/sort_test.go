package sort

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestSortInt(t *testing.T) {
	si := SliceFn[int]{
		s: []int{1, 3, 2, 5, 4},
		less: func(i, j int) bool {
			return i < j
		},
	}
	sort.Sort(si)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, si.s, "not equal")
}

func TestSortString(t *testing.T) {
	ss := SliceFn[string]{
		s: []string{"a", "ccc", "bb", "eeeee", "dddd"},
		less: func(i, j string) bool {
			return len(i) < len(j)
		},
	}
	sort.Sort(ss)
	assert.Equal(t, []string{"a", "bb", "ccc", "dddd", "eeeee"}, ss.s, "not equal")
}

func TestSortFn(t *testing.T) {
	si := []int{1, 3, 2, 5, 4}
	SortFn(si, func(i, j int) bool {
		return i < j
	})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, si, "not equal")
}
