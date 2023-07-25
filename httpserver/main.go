package main

import (
	"fmt"
	"net/http"
)

// 使用net/http包创建一个简单的web服务器
func main() {
	// 使用HandleFunc函数将handler注册到DefaultServeMux默认服务器路由,是net/http包的全局ServeMux实例
	// HandleFunc函数的第一个参数是路由的路径，第二个参数是处理该路由的函数
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 调用ListenAndServe函数，第一个参数是服务器监听的TCP地址，第二个参数是处理所有HTTP请求的处理器
	http.ListenAndServe(":8090", nil) // nil表示使用DefaultServeMux
}

// handlers是一个实现了http.Handler接口的对象
// http.ResponseWriter类型的参数被用于写入HTTP响应的数据
// http.Request类型的参数包含了所有的HTTP请求的信息
func hello(w http.ResponseWriter, req *http.Request) {
	// Fprintf函数将字符串写入到w中
	fmt.Fprintf(w, "hello\n")
}

// 读取HTTP请求的头信息，并将其写入到响应体中
func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}