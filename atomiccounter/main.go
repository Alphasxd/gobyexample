package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Go中最主要的状态管理机制是依靠通道间的通信来完成的
// 通常使用sync/atomic包中的原子计数器来提供对数值类型的安全访问
func main() {
	
	var ops uint64 // 无符号整型，用于存储计数器的值

	var wg sync.WaitGroup // WaitGroup用于等待这里启动的所有协程完成

	// 启动50个协程，并为其递增WaitGroup的计时器
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			// 每个协程将ops计数器递增1000次
			for c := 0; c < 1000; c++ {
				// 使用AddUint64来让计数器自动增加，使用&语法来给出ops的内存地址
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait() //等待，知道所有协程的完成

	// 安全的访问ops，此时不会有其他协程对其进行写操作
	// 此外还可以使用atomic.LoadUint64来拷贝一个安全的ops副本到opsFinal中
	fmt.Println("ops:", ops)
}