package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 将文件内容读取到内存中
	dat, err := os.ReadFile("file")
	check(err)
	fmt.Print(string(dat))

	// 对文件的读取方式和内容进行更多的控制
	f, err := os.Open("file")
	check(err)

	// 使用defer在完成后关闭这个文件
	defer f.Close()

	// 从文件的开始位置读取一些字节，最多允许读取5个字节，要注意实际读取了多少
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // 切片输出

	// Seek到一个文件中已知的位置，从这个位置开始读取
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// io包提供了一个更健壮的实现ReadAtLeast，用于读取上面的文件类型
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Seek(0, 0)实现`倒带`的功能
	_, err = f.Seek(0, 0)
	check(err)

	// bufio包实现了一个缓冲读取器，有助于提高许多小读操作的效率，以及提供了附加的读取函数
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5) // Peek返回缓冲区中的下n个字节，而不会移动读取位置
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}