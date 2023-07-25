package main

import (
	"fmt"
	"time"
)

// 速率限制是控制服务资源利用和质量的重要机制
// 基于协程、通道和Ticker，Go优雅的支持速率限制
func main() {
	
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 任务速率限制的调度器
	limiter := time.Tick(200 * time.Millisecond)

	// 在每次请求前阻塞limiter通道的接收，可以将频率限制为每200ms执行一次请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 通过缓冲通道来允许短暂的并发请求，并同时保留总体速率限制
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发 bursts
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func()  {
		// 每200ms添加一个新的值到burstyLimiter中，知道达到3个的限制
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 模拟另外5个传入请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}