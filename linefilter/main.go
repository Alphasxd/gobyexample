package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// 行过滤器（line filter）是一种常见的程序类型，
// 它读取 stdin 的输入，处理该输入，然后将得到的一些结果输出到 stdout。
func main() {

	// 使用带缓冲的scanner包装无缓冲的os.Stdin
    scanner := bufio.NewScanner(os.Stdin)

	// 可以将scanner看作是一个迭代器，每次调用scanner.Scan()就会读取下一行
    for scanner.Scan() {
		// Text()返回当前行的内容，不包括行尾标识符
		// ToUpper()将字符串转换为大写
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err) // Fprintln()将内容写入到指定的io.Writer
        os.Exit(1)
    }
}