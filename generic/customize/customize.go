/*
*
@link https://juejin.cn/post/7080938405449695268
*/
package main

// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
var a MyMap[string, float64] = map[string]float64{
	"jack_score": 9.6,
	"bob_score":  8.4,
}

var b MyMap[int, float64] = map[int]float64{
	1: 9.6,
	2: 8.4,
}

type Slice[T int | float32 | float64] []T

func (s Slice[T]) Len() int {
	return len(s)
}

type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

func main() {
}

/**
----------------------------------------
泛型 receiver
*/

type MySlice[T int | float64] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func (s MySlice[T]) Append(v T) MySlice[T] {
	return append(s, v)
}

/**
var m  = MySlice[int]{} // 泛型类型实例化

method 参数需定位为形参
*/

/**
----------------------------------------
泛型 queue, 结构体的定义和 method 的 receiver 都需要指定类型形参T
*/

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Put(v T) {
	q.elements = append(q.elements, v)
}

// 从队列头部取出并从头部删除对应数据
func (q *Queue[T]) Pop() (T, bool) {
	var value T
	if len(q.elements) == 0 {
		return value, true
	}
	value = q.elements[0]
	q.elements = q.elements[1:]
	return value, len(q.elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

/**
var q Queue[int]
q.Put(1)
q.Pop()

var q2 Queue[string]  // 可存放string类型数据的队列
q2.Put("A")
q2.Put("B")
q2.Put("C")
q2.Pop() // "A"
q2.Pop() // "B"
q2.Pop() // "C"

var q3 Queue[struct{Name string}]
var q4 Queue[[]int] // 可存放[]int切片的队列
var q5 Queue[chan int] // 可存放int通道的队列
var q6 Queue[io.Reader] // 可存放接口的队列

*/
