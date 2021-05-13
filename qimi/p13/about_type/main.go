package main

import "fmt"

// p13 自定义类型和类型别名

// 1. 自定义类型
type NewInt int

// 2. 类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T value:%v\n", a, a)
	fmt.Printf("type of b:%T value:%v\n", b, b)
}
