package main

import "fmt"

// 为什么需要接口？

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵~")
}

type dog struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

type person struct {
	name string
}

func (p person) say() {
	fmt.Println("嘤嘤嘤~")
}

// 在Go语言中接口（interface）是一种类型，一种抽象的类型。
// 接口不管是什么类型, 它只管要实现什么方法
// 一个抽象的类型, 只要实现了say()这个方法的类型就可以成为sayer类型
type sayer interface {
	say()
}

// 打函数
func da(da sayer) {
	da.say() // 不管传进来的是什么, 都要打. 打Ta, Ta就会叫, 就要执行Ta的say方法
}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{}
	da(p1)

	// s是sayer类型的接口
	var s sayer
	c2 := cat{name: "咪咪"}
	s = c2
	fmt.Println(s)
}
