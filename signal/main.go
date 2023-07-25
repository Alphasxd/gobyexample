package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 在计算机科学中，信号（英语：Signals）是Unix、类Unix以及其他POSIX兼容的操作系统中进程间通讯的一种有限制的方式。
// 它是一种异步的通知机制，用来提醒进程一个事件已经发生。
// 在Go中使用channel来处理信号
func main() {
	// 创建一个channel来接收信号，有缓冲区，可以缓存1个信号
	sigs := make(chan os.Signal, 1)

	// 注册sigs通道，使其成为可以接收特定信号的channel
	// signal.Notify(chan<- os.Signal, ...interface{})
	// SIGINT：程序终止(interrupt)信号，可以通过ctrl+C来发送
	// SIGTERM：程序结束(terminate)信号，kill命令的默认信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个channel来阻塞主进程，等待信号
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true // 通知主进程，信号已经处理完毕
	}()

	fmt.Println("awaiting signal")
	<-done // 阻塞主进程，等待信号处理完毕
	fmt.Println("exiting")
}