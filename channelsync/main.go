package main

import (
	"fmt"
	"time"
)

func main() {
	// 运行一个工作协程，并给予用于通知的通道
	done := make(chan bool)
	go worker(done)

	// 程序将在接收到通道中工作完成的通知前一直阻塞
	// <-done并不是返回一个值，而是从通道done中读取一个值
	<-done

	// 如果把 <- done 这行代码从程序中移除，程序甚至会在 worker 开始运行前就结束了
	
}

// 在协程中运行这个函数，done通道将被用于通知其他协程这个函数已经工作完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦
	done <- true
}