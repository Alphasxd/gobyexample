package main

import "fmt"

func main() {
	// 使用 range 来求 slice 或 array 中元素的和
	nums := []int{2, 3, 4}
	sum := 0

	// range 返回 index 和 value
	for index, num := range nums {
		fmt.Printf("index: %d, num: %d\n", index, num)
	}

	// 使用空白标识符忽略 index
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range在`map`中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// range只遍历`map`的`key
	for key := range kvs {
		fmt.Println("key:", key)
	}

	hellostr := "Hello, 世界"	
	
	// range在字符串中迭代`Unicode`码点
	for i, c := range hellostr {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
	fmt.Println(len([]rune(hellostr)))

	// 使用for循环遍历字符串
	// 中文字符使用UTF-8编码，是变长编码
	// 所以hellostr的长度是13，而不是9
	// 这个循环会打印出每个字符的UTF-8编码
	for i := 0; i<len(hellostr); i++ {
		fmt.Printf("index: %d, char: %c\n", i, hellostr[i])
	}
	fmt.Println(len(hellostr))
}