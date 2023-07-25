package main

import (
	"fmt"
	"time"
)

// 超时 对于一个需要连接外部资源，或者有耗时较长的操作的程序而言时很重要的
func main() {
	// 没有成功执行，超时实例
	c1 := make(chan string)

	// 假设执行一个外部调用，并在2秒后使用通道C1返回它的执行结果
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	// 如果这个调用超过了允许的1秒的时间，将会执行超时case，打印timeout 1，实现超时机制
	// time.After函数，它接受一个时间间隔参数，返回一个time.C类型的单向通道， 
	// time.C是一个 <-chan Time类型，也就是说，只能从这个通道中读取数据
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 成功接收实例，未超时
	c2 := make(chan string, 1)

	// 同样模拟一个外部调用操作，执行时间为2秒
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	// 超时阈值设置为3秒 
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
