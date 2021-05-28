package main

import "fmt"

// p14 结构体的初始化
/*
注意:
	结构体定义的时候,每行后面没有逗号
	结构体初始化的时候,每行后面都有逗号
*/
type person struct {
	name string
	age  int
	city string
}

func main() {
	// 1. 键值对初始化
	p1 := person{
		name: "小米",
		age:  20,
		city: "北京",
	}
	fmt.Printf("%#v\n", p1)

	// 2. 值的列表进行初始化
	p2 := person{
		"苹果",
		50,
		"California",
	}
	fmt.Printf("%#v\n", p2)
}
