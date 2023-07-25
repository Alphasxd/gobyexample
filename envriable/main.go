package main

import (
	"fmt"
	"os"
	"strings"
)

// 环境变量使一种向Unix程序传递配置信息的常见方式
func main() {
	// 使用Setenv来设置一个键值对，Setenv(key, value string) error
	os.Setenv("FOO", "1")
	// 使用Getenv获取一个键对应的值，如果键不存在，将会返回一个空字符串
	// Getenv(key string) string
	println("FOO:", os.Getenv("FOO"))
	println("BAR:", os.Getenv("BAR"))

	// 使用Environ来列出所有环境变量键值对，Environ() []string
	// 这个函数会返回一个KEY=value形式的字符串切片，可以使用strings.SplitN来得到键和值
	for _, e := range os.Environ() { // 遍历环境变量，只取value，忽略index索引
		// 以=为分隔符，分割成2个字符串，SplitN(s, sep string, n int) []string
		// SplitN的作用是将字符串s按照sep分割成n个字符串，返回一个字符串切片
		// 如果sep为空字符串，SplitN会将s切分成每一个unicode码值一个字符串
		// 如果n为0，返回nil；如果n大于0，返回最多n个子串；如果n小于0，返回所有的子串
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0]) // 只打印pair[0]，即key
	}
}