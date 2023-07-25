package main

import (
	"fmt"
	"sort"
)

// Go的sort包实现了用户自定义数据类型和内建数据类型的排序功能
func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs) //原地排序,直接改变给定的切片,而不是返回一个新切片
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 通过使用sort来检查一个切片是否时有序的
	isSorted := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", isSorted)
}