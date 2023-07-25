package main

import (
	"fmt"
	"os"
)

// defer用于确保程序在执行完成之前，会调用某个函数，通常用于资源清理、文件关闭等操作
// 对于需要自动释放的资源，可以使用 Go 语言中的 defer 语句来确保在函数返回前执行必要的清理操作
// defer的用途跟其他语言中的ensure或finally类似
func main() {
	f := createFile("defer.txt")
	// defer语句会将函数推迟到外层函数返回之前执行，即main函数的writeFile函数执行完后再执行closeFile函数
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close() // 关闭文件时，进行错误检查时非常重要的，即使再defer函数中也是如此

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}