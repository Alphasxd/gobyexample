package main

import (
	"fmt"
	"time"
)

// Go的select让你可以同时等待多个channel操作。
func main() {
	
	c1 := make(chan string)
	c2 := make(chan string)

	// 开始计时
	start := time.Now()

	// 从两个通道中选择
	go func() {
		// 各个通道将在一定时间后接受一个值，模拟并行的协程执行时造成的阻塞（例如RPC操作的耗时）
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 使用select关键字来同时等待这两个值，并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// 打印总耗时
	fmt.Println(time.Since(start))

	// 实际运行两秒左右，因为两个协程是并行执行的
}