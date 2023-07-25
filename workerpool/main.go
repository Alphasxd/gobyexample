package main

import (
	"fmt"
	"time"
)

// 使用协程与通道实现一个 工作池（worker pool）
func main() {

	const numJobs = 5

	// 为了使用worker工作池并且收集其结果，需要使用两个通道 
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动3个worker, 初始默认阻塞
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个jobs，然后close这些通道，表示任务完毕
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 同步等待所有任务完成
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// worker程序将在jobs频道上接收工作，并在results上发送相应的结果
// 每个worker会sleep一秒，以模拟一项昂贵（耗时一秒钟的）任务
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
