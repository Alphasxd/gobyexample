package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args 提供原始命令行参数访问功能。
	argsWithProg := os.Args
	// os.Args[1:] 保存程序全部的参数，因为第一个参数是程序路径，通过切片操作舍弃掉了
	argsWithoutProg := os.Args[1:]

	// 通过使用标准的索引位置方式，获取单个参数的值
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}