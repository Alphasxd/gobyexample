package main

import "fmt"

func main() {
	// 调用intSeq函数，将返回值赋给nextInt
	// 返回值是一个func() int类型的函数
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// 重新创建一个函数，重新赋值给nextInt
	newInts := intSeq()
	fmt.Println(newInts()) // 1

}

// intSeq函数返回一个在其函数体内定义的匿名函数
// 返回的函数使用闭包的方式隐藏变量i,返回的函数隐藏变量i以形成闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}