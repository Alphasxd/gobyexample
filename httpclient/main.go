package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// 使用Go中的net/http包来实现HTTP客户端
func main() {
	// 向服务端发送一个HTTP GET请求
	// http.get使创建http.Client对象并调用其Get方法的捷径
	// 使用了http.DefaultClient对象及其默认设置
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 关闭响应体，resp.Body存储了响应的主体内容

	fmt.Println("Response status:", resp.Status) // 打印响应状态码

	// NewScanner创建并返回一个从r读取数据的Scanner，NewScanner(r io.Reader) *Scanner
	scanner := bufio.NewScanner(resp.Body)
	// Scan()方法将Scanner定位到下一个标记，该标记由SplitFunc识别
	for i := 0; scanner.Scan() && i < 5; i++ { // 打印响应主体的前5行
		fmt.Println(scanner.Text())
	}
	// 检查Scan是否因为标记的长度超过缓冲区的容量而返回false
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}