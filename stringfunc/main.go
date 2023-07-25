package main

import (
    "fmt"
    "strings"
)

func main() {

	// Contains函数，判断字符串s是否包含substr
    fmt.Println("Contains:  ", strings.Contains("test", "es"))

	// Count函数，计算字符串s中有多少个不重复的substr，如果substr为空，则返回s的长度+1
    fmt.Println("Count:     ", strings.Count("test", "t"))

	// HasPrefix函数，判断字符串s是否以substr开头
    fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))

	// HasSuffix函数，判断字符串s是否以substr结尾
    fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))

	// Index函数，返回substr在字符串s中第一次出现的位置，如果没有出现，则返回-1
    fmt.Println("Index:     ", strings.Index("test", "e"))

	// Join函数，将字符串数组elems用字符串sep拼接起来
    fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))

	// Repeat函数，将字符串s重复count次
	// 如果 count 是负数，则函数会抛出一个 panic 异常。
	// 如果 len(s) * count 的结果溢出了，也会抛出一个 panic 异常
    fmt.Println("Repeat:    ", strings.Repeat("a", 5))

	// Replace函数，将字符串s中的old替换为new，n表示替换的次数，如果n<0，则全部替换
    fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
    fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))

	// Split函数，将字符串s用sep分割成字符串数组
    fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))

	// ToLower函数，将字符串s中的所有字符都转换为小写
    fmt.Println("ToLower:   ", strings.ToLower("TEST"))

	// ToUpper函数，将字符串s中的所有字符都转换为大写
    fmt.Println("ToUpper:   ", strings.ToUpper("test"))
    fmt.Println()

	// 获取字符串长度
    fmt.Println("Len: ", len("hello"))

	// 获取字符串中某个字符的ASCII码
    fmt.Println("Char:", "hello"[1])
}