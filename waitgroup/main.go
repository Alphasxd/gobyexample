package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WaitGroup用于等待这里启动的所有协程完成
	// 如果WaitGroup显式传递到函数，则应使用指针
	var wg sync.WaitGroup

	// 启动几个协程，并为其递增WaitGroup的计时器
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i //避免在每个协程闭包中重复利用相同的i值

		// 将worker的调用包装在一个闭包中，可以确保通知WaitGroup此工作线程已完成
		// 这样,worker线程本身就不必知道其执行的并发原语
		go func()  {
			defer wg.Done()
			worker(i)
		}()
	}

	// 阻塞，知道WaitGroup计数器恢复为0，即所有协程的工作都已经完成
	wg.Wait()
}

// 每个协程都会运行该函数
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) //睡眠一秒钟，以此来模拟耗时的任务
	fmt.Printf("Worker %d done\n", id)
}