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
	fmt.Printf("%s在跑...\n", p.name)
}

// 使用指针接收者实现接口: 只有类型指针能够保存到接口变量中
func (p *person) say() {
	fmt.Printf("%s在说...\n", p.name)
}

func main() {
	var m mover   // 定义一个mover类型的变量
	p1 := person{ // p1是person类型的值
		name: "小草莓",
		age:  18,
	}
	p2 := &person{ // p2是person类型的指针
		name: "香蕉君",
		age:  20,
	}
	m = p1
	m.move()

	var s sayer // 定义一个sayer类型的变量
	s = p2
	s.say()

	var a animal
	a = p2
	a.move()
	a.say()
}
