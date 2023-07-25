package main

import (
	"fmt"
	"time"
)

func main() {
	// 我们一般会这样 同步地 调用函数 f(s)
	f("direct")

	// 使用go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用
	go f("goroutine")

	// 也可以使用匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 两个协程在独立的 Go 协程中异步的运行，所以我们需要使用time.Sleep来等待它们执行结束
	// 更好的方法是使用 WaitGroup 的计数信号来等待两个 Go 协程的结束
	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}