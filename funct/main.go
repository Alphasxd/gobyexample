package main

import "fmt"

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// 如果有一个多值的 slice，想把它们作为参数使用，需要这样调用 func(slice...)。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}