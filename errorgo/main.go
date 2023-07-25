package main

import (
	"errors"
	"fmt"
)

// 下面的两个循环测试了每一个会返回错误的函数
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 类型断言来得到自定义错误类型的实例
	// 类型断言的语法：i.(T)  i表示接口变量，T表示类型
	
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

// 错误通常是最后一个返回值并且是error类型，一个内建的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// 返回错误值为nil表示没有错误
	return arg + 3, nil
}

// 通过实现Error()方法来自定义error类型
type argError struct {
	arg int
	prob string
}

func (e * argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 使用&argError语法来建立一个新的结构体,并提供了arg和prob两个字段的值
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}