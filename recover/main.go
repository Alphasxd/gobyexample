package main

import "fmt"

// Go通过再defer中使用recover函数，可以截获panic异常，并返回一个非空的error，从而恢复正常的程序运行
// 注意recover函数只有在defer函数中才有效，否则无法截获panic异常
// 并且recover调用需放在defer中的第一行，否则无法截获panic异常
// 如果程序正常运行，调用recover会返回nil，且没有其他效果
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// 这行代码并不会执行
	// main函数在panic点停止，并在继续处理完defer后结束，recover函数只是保证了不将错误信息打印出来而已
	fmt.Println("After mayPanic()") 
}

func mayPanic() {
	panic("a problem")
}