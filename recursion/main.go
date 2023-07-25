package main

import "fmt"

func main() {
	// 7! = 7 * 6 * 5 * 4 * 3 * 2 * 1 = 5040
	fmt.Println(fact(7))

	// 显式声明一个函数变量fib
	var fib func(n int) int

	// 闭包递归函数，需要提前用var声明
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// 因为提前声明了fib函数，所以可以直接调用
		return fib(n-1) + fib(n-2)
	}

	// fibbonacci数列 1 1 2 3 5 8 13
	fmt.Printf("fib(7): %v\n", fib(7))
}

// fact函数在到达fact(0)之前一直调用自身
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}