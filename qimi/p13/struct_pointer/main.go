package main

import "fmt"

// 结构体指针

type person struct {
	name string
	age  int
	city string
}

func main() {
	var p1 = new(person)
	fmt.Printf("Type p1: %T\n", p1)
	(*p1).name = "小草莓" // 标准写法
	p1.age = 26        // 语法糖
	p1.city = "上海"     // 语法糖
	fmt.Printf("%#v\n", p1)

	// 取结构体的地址进行实例化
	p2 := &person{}
	fmt.Printf("Type p2: %T\n", p2)
	p2.name = "大草莓"
	p2.age = 31
	p2.city = "旧金山"
	fmt.Printf("%#v\n", p2)
}
