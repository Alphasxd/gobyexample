package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 非空的通道关闭后，通道中剩下的值仍然可以被接收
	close(queue)
	// 如果不关闭通道，会出现死锁的错误，deadlock

	// range 循环会在通道中的所有值都被接收后自动结束
	for elem := range queue {
		fmt.Println(elem)
	}
}