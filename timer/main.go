package main

import (
	"fmt"
	"time"
)

// 我们经常需要在未来的某个时刻执行 Go 代码，或者在某段时间间隔内重复执行。
// Go 的内置 定时器 和 打点器 特性让这些很容易实现。
func main() {
	// 定时器表示在未来某一时刻的独立事件。你告诉定时器需要等待的时间，然后它将提供一个用于通知的通道。
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C 将一直阻塞, 直到这个定时器的通道 C 明确的发送了定时器失效的值
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 如果仅仅是单纯的等待，只需要使用 time.Sleep。
	// 定时器是有用原因之一就是你可以在定时器失效之前，取消这个定时器。
	timer2 := time.NewTimer(time.Second)
	go func()  {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 给timer2足够的时间来触发，以证明它实际上已经停止了
	time.Sleep(2 * time.Second)
}