package main

import "fmt"

// 使用一个带default子句的select来实现非阻塞的发送和接受，甚至是非阻塞的多路select
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 非阻塞接收
	select {
	// 这里的msg := <-messages 尝试使用通道接收值，这里将会阻塞，因为没有任何协程向通道发送数据
	case msg := <-messages:
		fmt.Println("received message:", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送
	msg := "hi"
	select {
	// 这里的msg := <-messages 尝试使用通道发送值，这里将会阻塞，因为没有任何协程从通道接收数据
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default:
		fmt.Println("no message sent")
	}

	// 多路非阻塞select
	select {
	case msg := <-messages:
		fmt.Println("received message:", msg)
	case sig := <-signals:
		fmt.Println("received signal:", sig)
	default:
		fmt.Println("no activity")
	}
}