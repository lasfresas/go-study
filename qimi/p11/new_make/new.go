package main

import "fmt"

// new函数
/*
new是一个内置的函数，它的函数签名如下：
	func new(Type) *Type
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
*/

func main() {
	var a *int
	a = new(int)
	var b *bool
	b = new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}
