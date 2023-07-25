package main

import "fmt"

func main() {
	// 使用 make(chan val-type) 创建一个新的通道。通道类型就是他们需要传递值的类型
	messages := make(chan string)

	go func() {
		// 使用 channel <- 语法 发送(send) 一个新的值到通道中
		messages <- "ping"
	}()

	// 使用 <-channel 语法从通道中 接收(receives) 一个值
	if msg := <-messages; msg == "ping" {
		fmt.Println("pong")
	}

	// 默认情况下通道是无缓冲的，所以通道的发送和接受操作是阻塞的，直到发送方和接收方都准备完毕。
	// 所以在上面的例子中，程序会一直等待ping，直到它被接收，所以这里不用额外的等待时间和其他的同步操作
}