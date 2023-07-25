package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex // 使用互斥锁来使多个goroutine同步访问counters这个map
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock() // 访问counter之前锁定
	defer c.mu.Unlock() // 函数结束时解锁
	c.counters[name]++
}

// 使用互斥量在Go协程间安全地访问数据
func main() {
	// 互斥量有初始零值，因此不需要初始化
	c := Container {
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait() // 等待所有goroutine完成
	fmt.Println(c.counters)
}