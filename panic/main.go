package main

import "os"

// Panic意味着有些出乎意料的错误发生。
// 通常用来表示程序正常运行中不应该出现的，或者不准备优雅处理的错误。
func main() {

	// panic 清理(panic cleanup)是指在程序发生 panic 异常时，依次执行所有的 defer 语句
	// 以确保程序能够正确地清理资源和执行其他必要的操作。
	defer panic("a problem")

	// panic("a problem") 语句会导致程序崩溃并抛出一个 panic 异常。
	// 因此，这行代码之后的所有代码都不会被执行
	// 正确做法是将第8行前加defer, 这样即使出错误，也会在执行完其他代码后再执行defer中的代码
	_, err := os.Create("/tmp/file") // unreachable code

	if err != nil {
		panic(err)
	}
}