package main

import "fmt"

func main() {
	// 创建一个空 map，使用内建的 make: make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	// 打印 map
	fmt.Println(m)

	// 内建函数 len 返回 map 的长度
	fmt.Println("length of map:", len(m))

	// 如果 key 存在，返回对应的 value 和 true
	// 如果 key 不存在，返回对应类型的零值和 false
	v, ok := m["k3"]
	fmt.Println("value:", v)
	fmt.Println("key exists:", ok)

	// 使用空白标识符忽略 value
	_, ok = m["k2"]
	fmt.Println("only key exists:", ok)

	delete(m, "k2")
	fmt.Println("map after delete k2:", m)

	// 声明并初始化 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("declare and init map:", n)
}