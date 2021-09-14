package bitmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestBitMap_Add(t *testing.T) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("memeory start cost: ", ms.Sys>>20)
	max := 20000000
	b := NewBitMap()
	tests := []struct {
		name   string
		num    int
		length int
	}{
		{name: "add 1", num: 1, length: 1},
		{name: "add 2", num: 2, length: 2},
		{name: "add 3", num: 3, length: 3},
		{name: "add 1 again", num: 1, length: 3},
		{name: "add 20000000", num: max, length: 4},
		{name: "add 0", num: 0, length: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.Add(tt.num)
			if b.length != tt.length {
				t.Errorf("%s, expect length = %d, actual = %d", tt.name, tt.length, b.length)
			}
		})
	}
	wordsLength := max / 64
	wordsLength += 1
	t.Run("check words length", func(t *testing.T) {
		assert.Equal(t, wordsLength, len(b.words), "expect words length = %d, actual = %d", wordsLength, len(b.words))
	})
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("memeory end cost: ", m.Sys>>20)
}

func TestBitMap_Has(t *testing.T) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("memeory start cost: ", ms.Sys>>20)
	b := NewBitMap()
	nums := []int{1, 2, 3, 4, 10000000, 14845973, 14845971}
	for _, v := range nums {
		b.Add(v)
	}

	tests := []struct {
		name string
		num  int
		want bool
	}{
		{name: "check 1", num: 1, want: true},
		{name: "check 2", num: 2, want: true},
		{name: "check 3", num: 3, want: true},
		{name: "check 4", num: 4, want: true},
		{name: "check 5", num: 5, want: false},
		{name: "check 6", num: 6, want: false},
		{name: "check 7", num: 7, want: false},
		{name: "check 10000000", num: 10000000, want: true},
		{name: "check 14845973", num: 14845973, want: true},
		{name: "check 14845971", num: 14845971, want: true},
		{name: "check 14845966", num: 14845966, want: false},
		{name: "check 14845956", num: 14845956, want: false},
		{name: "check 0", num: 0, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := b.Has(tt.num)
			assert.Equal(t, tt.want, got, "Has(%d) = %v, want %v", tt.num, got, tt.want)
		})
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("memeory end cost: ", m.Sys>>20, ", nums words: ", len(b.words))
}

func TestBitMap_Remove(t *testing.T) {
	b := NewBitMap()

	b.Add(100)
	got := b.Has(100)
	if got != true {
		t.Errorf("check 100 fail: expect %v, actual %v", got, true)
	}
	b.Remove(100)
	got = b.Has(100)
	if got != false {
		t.Errorf("check 100 fail: expect %v, actual %v", got, false)
	}
}
