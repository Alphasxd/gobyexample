package main

import "fmt"

type rect struct {
	width, height int
}

// GO支持为结构体类型定义方法
// 方法接收者出现在func关键字和方法名之间的参数中
// 这里的area是一个拥有*rect类型接收器的方法
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法
// 这是一个值类型接收器的例子 (仅供示例使用)
// 依据 Go 的约定，如果结构体的任何方法具有指针接收方，
// 则此结构的所有方法都必须具有指针接收方。 
// 即使此结构的某个方法不需要它也是如此
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perim())

	rptr := &r
	// 因为结构体指针变量会自动解引用，所以可以直接调用方法
	fmt.Println("area:", rptr.area())
	fmt.Println("perimeter:", rptr.perim())
}