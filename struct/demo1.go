package main

import "fmt"

type A struct {
	Age int
}

func (a A) Say() {
}

type B struct {
}

func (b B) Say() {}

type ABer interface {
	Say()
}

func main() {
	a1 := &A{Age: 1}
	a2 := &A{Age: 1}
	a3 := &A{Age: 1}
	a4 := a3
	fmt.Println(a1 == a2) //指针比较为false, 值比较为true
	fmt.Println(a1.Age == a2.Age)
	fmt.Println(a3 == a4)
	var ab1 ABer = A{}
	var ab2 ABer = B{}
	var ab3 ABer = B{}
	fmt.Println(ab1 == ab2) //false,接口比较时类型、值相等才算相等
	fmt.Println(ab2 == ab3) //true
}
