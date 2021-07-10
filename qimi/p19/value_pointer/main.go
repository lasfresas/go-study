package main

import "fmt"

// 使用值接收者实现接口和使用指针接收者实现接口的区别

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int
}

// 使用值接收者实现接口: 类型的值和类型的指针都能保存到接口变量中
func (p person) move() {
	fmt.Printf("%s在跑...\n", p.name) // 都能跑
}

// 使用指针接收者实现接口: 只有类型指针能够保存到接口变量中
func (p *person) say() {
	fmt.Printf("%s在说...\n", p.name) // 但只有指针类型的能说
}

func main() {
	p1 := person{ // p1是person类型的值
		name: "小草莓",
		age:  20,
	}
	var m mover // 定义一个mover类型的变量
	m = p1
	m.move()

	p2 := &person{ // p2是person类型的指针
		name: "香蕉君",
		age:  30,
	}
	var s sayer // 定义一个sayer类型的变量
	s = p2
	s.say()

	p3 := &person{ // p3是person类型的指针
		name: "大雕怪",
		age:  15,
	}
	var a animal // 定义一个animal类型的变量
	a = p3
	a.move()
	a.say()
	fmt.Println(a)
}
