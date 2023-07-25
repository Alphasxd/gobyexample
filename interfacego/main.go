package main

import (
	"fmt"
	"math"
)

// 定义一个几何体接口
type geometry interface {
	area() float64
	perim() float64
}

// 定义一个矩形结构体
type rect struct {
	width, height float64
}

// 定义一个圆形结构体
type circle struct {
	redius float64
}

// rect实现geometry接口
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * (r.width + r.height)
}

// circle实现geometry接口
func (c circle) area() float64 {
	return c.redius * c.redius * math.Pi
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

// 如果一个变量实现了某个接口，那么就可以调用指定接口中的方法
func measure(g geometry) {
	fmt.Printf("%T\n", g)
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	r := rect{3, 4}
	c := circle{5}

	// 因为结构体rect和circle都实现了geometry接口
	// 所以可以将其实例作为参数传递给measure函数
	measure(r)
	measure(c)
}