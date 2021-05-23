package main

import "fmt"

// 值接收者和指针接收者

type person struct {
	name string
	age  int
}

// 值接收者：传拷贝进去
func (p person) guonian() {
	p.age++
}

// 指针接收者：传内存地址进去
func (p *person) zhenguonian() {
	p.age++
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("小明", 18)
	fmt.Println(p1.age) // 18
	p1.guonian()
	fmt.Println(p1.age) // 18
	p1.zhenguonian()
	fmt.Println(p1.age) // 19
}
