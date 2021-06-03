package main

import "fmt"

// 使用指针接收者实现接口,只能存结构体的指针类型的变量

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int
}

// 使用指针接收者实现了接口的所有方法
func (c *cat) move() {
	fmt.Println("走猫步!")
}
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s!\n", food)
}

func main() {
	var a1 animal
	var a2 animal

	c1 := cat{"tom", 4}  // cat
	c2 := &cat{"假老练", 4} // *cat

	a1 = &c1 // 实现animal这个接口的是cat的指针类型
	fmt.Println(a1)
	a2 = c2
	fmt.Println(a2)
}
