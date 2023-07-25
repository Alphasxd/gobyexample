package main

import "fmt"

func main() {
	// 使用 make() 创建空 slice
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	fmt.Println("length of slice:", len(s))

	// append slice
	// append() 函数会返回一个包含了一个或者多个新值的 slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	// copy slice
	// copy() 函数将 src 中的内容复制到 dst 中，返回复制的元素个数
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copyed slice:", c)

	// slice operator
	// syntax: slice[low:high] 左闭右开区间
	l := s[2:5]
	fmt.Println("slice[2:5]:", l)

	l1 := s[:5]
	l2 := s[2:]

	fmt.Println("slice[:5]:", l1)
	fmt.Println("slice[2:]:", l2)

	// 声明并初始化 slice
	t := []string{"g", "h", "i"}
	fmt.Println("declare and init slice:", t)

	// multi-dimension slice
	// 内部的 slice 长度可以不一致，这一点和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("multi-dimension slice:", twoD)
}
