package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 结构体readOp和writeOp封装了对协程的请求，分别用于读和写
type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	// 协程通过reads和writes通道来发布读和写请求
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 拥有state的协程，state被协程私有
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 100个协程通过reads通道向拥有state的协程发起读取请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 10个协程通过writes通道向拥有state的协程发起写入请求
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 让协程跑1s
	time.Sleep(time.Second)

	// 获取并报告ops
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}