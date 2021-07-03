package main

import "fmt"

/*
方法只能写在类型所在的包里,不能给别的包的类型定义方法,类型和类型的方法必须在一个包
*/

// MyInt 基于int类型自定义MyInt类型
type MyInt int

// SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个MyInt。")
}

func main() {
	var m1 MyInt
	m1.SayHello() // Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v %T\n", m1, m1) // 100 main.MyInt
}
