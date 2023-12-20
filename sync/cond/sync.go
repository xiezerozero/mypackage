package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		condition bool
		c         = sync.NewCond(&sync.Mutex{})
	)

	for i := 0; i < 7; i++ {
		go func(index int) {
			c.L.Lock()
			for !condition {
				c.Wait()
			}
			fmt.Println("index:", index)
			c.L.Unlock()
			// ...
		}(i)
	}
	// 保证让其他7个协程进入wait状态
	time.Sleep(time.Millisecond)
	// Elsewhere
	c.L.Lock()
	condition = true
	c.L.Unlock()
	// 修改条件, 广播事件, 让其他挂起的协程继续执行
	c.Broadcast()
	fmt.Println("broadcast")

	time.Sleep(time.Second)
}
