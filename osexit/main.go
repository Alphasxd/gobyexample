package main

import (
	"fmt"
	"os"
)

// 使用os.Exit()方法退出程序
// Exit()的状态代码有一个惯例，0表示成功，非0表示错误
func main() {
	// defer不会被执行，因为os.Exit()会立即退出程序
	// 当程序调用os.Exit函数时，程序会立即终止，不会运行defer()函数
	defer fmt.Println("!")
	os.Exit(3)
}