package generic

import (
	"fmt"
	"sync"
)

// constraints on type parameters
type Number interface {
	// ~int 代表基础类型为int的都属于Number
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		~float32 | ~float64
}

func Max[T Number](a ...T) T {
	var max T
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

type Age int //Age底层是int, ~int 满足,可直接使用Age类型调用Add方法,如果未定义~int, Age类型调用编译不过

func Add[T Number](a, b T) T {
	return a + b
}

// 针对确定的类型进行联合约束
func PrintInt64OrFloat64[T int64 | float64](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

type LazyLoad[T any] struct {
	dataMap sync.Map
	f       func(string) T
}

func NewLazyLoad[T any](f func(key string) T) *LazyLoad[T] {
	return &LazyLoad[T]{
		f: f,
	}
}

type LazyData[T any] struct {
	data T
}

// bool, 代表是否是从syncMap中获取
func (l *LazyLoad[T]) Get(key string) (bool, T) {
	if v, isLoaded := l.dataMap.Load(key); isLoaded {
		innerData, ok := v.(LazyData[T])
		if ok {
			return true, innerData.data
		} else { //类型不对,重新加载一次
			l.dataMap.Delete(key)
		}
	}
	if l.f != nil {
		data := l.f(key)
		v := LazyData[T]{data: data}
		l.dataMap.Store(key, v)
		return false, data
	}
	return false, l.Zero()
}

func (l *LazyLoad[T]) Set(key string, data LazyData[T]) {
	l.dataMap.Store(key, data)
}

func (l *LazyLoad[T]) Del(key string) {
	l.dataMap.Delete(key)
}

// refresh data by set dataMap = sync.Map{}
func (l *LazyLoad[T]) Refresh() {
	l.dataMap = sync.Map{}
}

func (l *LazyLoad[T]) Zero() T {
	var zero T
	return zero
}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// Add method on Set
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Intersect method on set, accept another set as param
func (s Set[T]) Intersect(t Set[T]) Set[T] {
	var result = Set[T]{}
	for k := range s {
		if _, ok := t[k]; ok {
			result.Add(k)
		}
	}
	return result
}

// Union method on set, accept another set as param
func (s Set[T]) Union(t Set[T]) Set[T] {
	var result = Set[T]{}
	for k := range s {
		result.Add(k)
	}
	for k := range t {
		result.Add(k)
	}
	return result
}

// IndexFunc returns the first index i satisfying f(s[i]),
// or -1 if none do.
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := range s {
		if f(s[i]) {
			return i
		}
	}
	return -1
}
