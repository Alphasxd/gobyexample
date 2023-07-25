package main

import "fmt"

// GO的结构体是带类型的字段集合，这在组织数据时非常有用。
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// 创建新的结构体元素并打印
	fmt.Println(person{"Bob", 20})
	// 初始化结构体元素时指定字段名字
	fmt.Println(person{name: "Alice", age: 30})
	// 省略的字段将被初始化为零值
	fmt.Println(person{name: "Fred"})
	// & 前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})
	// 封装新建结构体的函数
	fmt.Println(newPerson("John"))

	// 使用`.`来访问结构体字段
	s := person{"Sean", 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	// 使用结构体指针来访问结构体字段
	// 指针会被自动解引用，无需使用`*`操作符，即(*sptr).name
	sptr := &s
	fmt.Println(sptr.name)

	// 结构体是可变的
	sptr.age = 51
	fmt.Println(s.age)

}