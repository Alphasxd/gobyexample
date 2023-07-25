package main

import "fmt"

// Go支持对于结构体和接口的嵌入，以表达一种更加无缝的组合类型
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base // 一个containner嵌入了一个base， 是一个没有名字的字段field
	str string
}

func main() {
	ctn := container{
		base: base{1},
		str: "some name",
	}

	// 可以直接在ctn上访问base中定义的字段和方法
	fmt.Printf("ctn = {num: %v, str: %v}\n", ctn.num, ctn.str)
	fmt.Println(ctn.describe())

	// 使用完整路径访问嵌入的字段和方法
	fmt.Println("also num:", ctn.base.num)
	fmt.Println("also describe:", ctn.base.describe())

	type describer interface {
		describe() string
	}
	// 使用带有方法的嵌入结构 来赋予接口实现 到其他结构上
	// 因为container嵌入了base结构, base结构有describe()方法
	// 因此containner也实现了describer接口
	var d describer = ctn
	fmt.Println("describer:", d.describe())
}