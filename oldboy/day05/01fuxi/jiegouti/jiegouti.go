package main

import "fmt"

// 复习结构体

type tmp struct {
	x int
	y int
}

type person struct {
	name string
	age  int
}

// 构造(结构体变量的)函数，返回值是对应的结构体类型
func newPerson(n string, i int) person {
	return person{
		name: n,
		age:  i,
	}
}

func main() {
	// 匿名结构体, 多用于临时场景
	var a = struct {
		x int
		y int
	}{10, 20}
	fmt.Println(a)

	var b = tmp{
		10,
		20,
	}
	fmt.Println(b)

	type person struct {
		name string
		age  int
	}

	// 声明结构体
	// 结构体实例化
	var p1 person
	p1.name = "小明"
	p1.age = 20

	// 结构体实例化
	p2 := person{"小红", 30}

	// 调用构造函数生成person类型变量
	p3 := newPerson("小习", 10000)

	fmt.Println(p1, p2, p3)
}
