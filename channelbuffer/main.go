package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	// 因为通道的缓冲区大小为2，所以我们可以将这些值发送到通道中，而不需要并发的接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以像前面一样接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}