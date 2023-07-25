package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256经常用于生成二进制文件或者文本块的短标识
// 例如：git就是用SHA256来标识受版本控制的文件和目录
// go.sum也是用SHA256来标识依赖包的版本
func main() {
	s := "sha256 this string"

	h := sha256.New()

	h.Write([]byte(s)) // 强制转换为byte切片

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}