package main

import (
	"fmt"
	"net/http"
	"time"
)

// context.Context对于HTTP请求来说是一个很好的选择，
// 因为它不仅仅在请求的处理函数中传递请求的数据，还能在请求的整个生命周期中传递请求的数据
// Context跨API边界和协程边界传递请求的数据，例如：deadline、取消信号和其他请求范围的值
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	// 通过req.Context()方法创建一个Context对象，这个Context对象可以在请求的整个生命周期中传递请求的数据
	// func(*Request) Context() context.Context
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	// 通过time.After()方法创建一个10秒的定时器，返回一个<-chan Time类型的channel
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err() // 返回一个错误类型的值，说明了Done通道为什么关闭
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError // 错误码500
		http.Error(w, err.Error(), internalError) // 将错误信息写入到响应体中
		
	}
}

