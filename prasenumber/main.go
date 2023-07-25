package main

import (
	"fmt"
	"strconv"
)

// 使用Go来解析字符串，可以使用内置的strconv包
func main() {
	// 使用PraseFloat函数将字符串转换为浮点数，这里的64表示转换为float64，返回值的第二个参数表示转换的精度
	f, _ := strconv.ParseFloat("1.234", 64) 
	fmt.Printf("%T: %v\n", f, f)

	// int类型的位数为32位或64位，具体取决于平台的位数。因此，如果您不确定要使用哪种整数类型
	// 可以将bitSize参数设置为0，并让ParseInt函数自动选择一个与当前平台相同位数的整数类型
	// 当ParseInt函数的第二个参数为0时，表示自动推断字符串所表示的数字的进制
	i, _ := strconv.ParseInt("212", 0, 64)
	fmt.Printf("%T: %v\n", i, i)

	// ParseInt函数还可以解析16进制的字符串
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Printf("%T: %v\n", d, d)

	// ParseUint函数可以解析无符号整型
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Printf("%T: %v\n", u, u)

	// Atoi函数等价于ParseInt(s, 10, 0)，用于转换字符串为int类型，此处的0表示自动推断字符串所表示的数字的进制
	// Atoi 函数的参数必须为“整数型”的字符串，否则会返回错误
	k, _ := strconv.Atoi("135")
	fmt.Printf("%T: %v\n", k, k)

	_, e := strconv.Atoi("wait")
	fmt.Println(e)

}