package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	// 定义命令行参数
	var name string
	var age int
	var married bool
	var delay time.Duration

	// 此种方式变量是已经声明好的，需要在命令行中指定参数值，更加适用于结构体中的字段
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "delay", 0, "延迟时间间隔") // 默认0代表不延迟

	// 另外一种定义命令行参数的方式，这种方式返回的是指针
	// name := flag.String("name", "张三", "姓名")
	// age := flag.Int("age", 18, "年龄")
	// married := flag.Bool("married", false, "婚否")
	// delay := flag.Duration("delay", 0, "延迟时间间隔")

	// 解析命令行参数，从os.Args[1:]中解析命令行参数，默认是从第一个参数开始解析
	flag.Parse()

	// 解析命令行参数
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("married:", married)
	fmt.Println("delay:", delay)

	// 对应另一种定义命令行参数的方式，这种方式返回的是指针，所以要对指针进行取值操作
	// fmt.Printf("%v %v %v %v\n", *name, *age, *married, *delay)


	// 返回命令行参数后的其他参数
	fmt.Println("Tail slice ", flag.Args())

	// 返回命令行参数后的其他参数数量
	fmt.Println("Tail slice length ", flag.NArg())

	// 返回使用的命令行参数数量
	fmt.Println("Used flags count", flag.NFlag())
}