package main

import "fmt"

func main() {
	i := 1
	fmt.Println("initial:", i) // 1

	zeroval(i)
	fmt.Println("zeroval:", i) // 1

	// &操作符会生成一个指向其操作数的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0

	fmt.Println("pointer:", &i)
}

// zeroval有一个int参数，所以使用值传递
// zeroval将从调用它的那个函数中得到一个实参的拷贝：ival(形参)
func zeroval(ival int) {
	// ival = 0
}

// zeroptr有一个*int参数，所以使用指针传递
// 调用方函数实参的地址值被拷贝到了形参中，实际上也是值传递
func zeroptr(iptr *int) {
	// *操作符解引用指针，从它的内存地址得到这个地址当前对应的值
	// 对解引用的指针赋值，会改变这个指针引用的真实地址的值
	*iptr = 0
}