package main

import (
	"fmt"
	"time"
)

// 定时器是当你想要在未来某一刻`执行一次`时使用的
// 打点器则是当你想要在固定的时间间隔`重复执行`准备的
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func()  {
		for {
			select {
			// done通道在接收到值之前，会一直阻塞，一旦接收到值，就会立即返回并停止循环
			case <-done:
				return
			// ticker.C 是一个通道，每隔 500ms 就会发送一次
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Ticker可以和Timer一样被停止
	// Ticker一旦被停止，将不能再从它的通道中接收到值
	time.Sleep(1600 * time.Millisecond)

	// ticker的stop()没有返回类型，与timer不同 
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}