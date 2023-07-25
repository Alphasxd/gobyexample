package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go的math/rand包提供了`伪随机数`生成器 pseudo-random number generators
func main() {
	fmt.Print(rand.Intn(100), ", ") // Intn返回一个随机的整数n，0 <= n < 100，即[0, 100)
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64()) // Float64返回一个64位浮点数f，0.0 <= f < 1.0，即[0, 1)

	fmt.Print((rand.Float64()*5)+5, ", ") // 可以用来生成其他范围的随机浮点数，此案例即[5, 10)
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()

	// 默认情况下，给定的种子是确定的，每次都会产生相同的随机数数字序列
	// 要产生变化的序列，需要给定一个变化的种子
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ", ")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ", ")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ", ")
	fmt.Print(r3.Intn(100))
	fmt.Println()

}