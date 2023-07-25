package main

import "fmt"

// 关闭一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
func main() {

	// 使用一个jobs通道，将工作内容，从main()线程传递到一个工作协程中。没有多余的工作时，将关闭jobs通道。
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是工作协程。使用 j, more := <- jobs 循环的从 jobs 接收数据。 
	// 根据接收的第二个值，如果 jobs 已经关闭了， 并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。 
	// 当完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。
	go func() {
		for {
			// 使用两个返回值来检测通道是否关闭，如果没有值可以接收，并且通道已经关闭，那么more的值将是false
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	
	// 使用jobs通道发送三个任务到工作协程中，然后关闭jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// close函数只能用来关闭发送通道或者双向通道，不能用来关闭接收通道。
	// 向一个已经关闭的通道发送数据会引起panic，向一个已经关闭的通道接收数据是安全的。
	close(jobs)
	fmt.Println("sent all jobs")

	// 通道同步方法，等待工作协程完成
	<- done

}