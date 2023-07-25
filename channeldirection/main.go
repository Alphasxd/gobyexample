package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// 当使用通道作为函数的参数时，你可以指定这个通道是不是只用来发送或者接收值
// 这个特性提升了程序的类型安全性, 防止编写失误

// ping 函数定义了一个只允许发送数据的通道（只写），尝试使用这个通道来接收数据将会在编译时报错
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接收两个通道，pings 仅用于接收数据（只读），pongs 仅用于发送数据（只写）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}