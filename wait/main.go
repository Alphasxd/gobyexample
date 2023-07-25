package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "🐏"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "🐂"
			time.Sleep(time.Millisecond * 2000)
		}
	}()

	// 如果不使用select，那么打印的结果将是两个通道的值交替出现，
	// 并不会根据时间间隔来打印，因为阻塞了两个通道的接收操作。
	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
