package main

import (
	"fmt"
	"sort"
)

// 使用自定义函数排序需要对应类型
// 因为不能对内建类型创建方法（method），所以这里创建一个内建类型[]string的别名
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 自定义排序示例，例如按照长度排序而不是按照alphabetical排序
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits)) // 将fruits转换为byLength类型
	fmt.Println(fruits)
}