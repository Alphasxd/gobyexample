package main

import (
	"fmt"
	"unicode/utf8"
)

// Go语言中的字符串是一个只读的byte类型的切片。
func main() {
	const s = "你好世界"

	// 因为字符串等价于[]byte, 所以len(s)返回原始字节的长度
	fmt.Println("length of s:", len(s)) // 6

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 使用utf8.RuneCountInString(s)获取字符串中的rune数量
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 使用utf8.DecodeRuneInString函数实现对单个字符的解码
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

	// range循环专门处理字符串并解码每个rune及其字符串中的偏移量
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, idx)
	}
}

func examineRune(r rune) {
	if r == 'u' {
		fmt.Println("found you")
	} else if r == '你' {
		fmt.Println("found 你")
	}
}
