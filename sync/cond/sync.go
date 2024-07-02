package main

import (
	"fmt"
)

func main() {
	a := 3.1232
	fmt.Println(fmt.Sprintf("失效次数超过%.0f次,请及时处理", a))
}
