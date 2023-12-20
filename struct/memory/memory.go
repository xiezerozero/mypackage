/**
https://prog-bytes.hashnode.dev/golang-structs-memory-allocation
golang 结构体内存分配

各类型占用的字节数, 为了内存对齐,每个类型的字段的位置需要对齐
eg: string类型内存起始位置  needs to be divisible by 8

byte, int8, uint8 -> 1
int16, uint16 -> 2
int32, uint32, float32, complex64 -> 4
int, int64, uint64, float64, complex128 -> 8

string, slice -> 8
string是使用stringHeader 存储,实际占用16个字节
slice使用sliceHeader 存储,实际占用24个字节
*/
package main

import (
	"fmt"
	"unsafe"
)

// 如果将b放在最后面, a,c,d 就会占用前8个字节: a 1个字节 c一个字节 d4个字节
type Example struct {
	a int8
	c int8
	d int32
	b string
}

func main() {
	var v = Example{
		a: 10,
		b: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus rhoncus.",
		c: 20,
		d: 100,
	}
	fmt.Println("Offset of a", unsafe.Offsetof(v.a), ", SizeOf a ", unsafe.Sizeof(v.a)) // 0
	fmt.Println("Offset of b", unsafe.Offsetof(v.b), ", SizeOf b ", unsafe.Sizeof(v.b)) // 8
	fmt.Println("Offset of c", unsafe.Offsetof(v.c), ", SizeOf c ", unsafe.Sizeof(v.c)) // 24
	fmt.Println("Offset of d", unsafe.Offsetof(v.d), ", SizeOf d ", unsafe.Sizeof(v.d)) // 28

	fmt.Println(unsafe.Sizeof(v))
}
