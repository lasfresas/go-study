package main

import "fmt"

// p15 方法和接收者

// Person 是一个结构体
type Person struct {
	name string
	age  int
}

// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name, // 前面的name是结构体的name字段,后面的name是传入函数的参数
		age:  age,
	}
}

// Dream 是为Person类型定义方法。
// 值接收者
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

// SetAge 是一个修改年龄的方法。
// 指针接收者
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

func main() {
	p1 := NewPerson("小草莓", 20)
	p1.Dream()
	p1.SetAge(9000)
	fmt.Println(p1.age)
}
