package main

import "fmt"

// 自定义类型和类型别名

// MyInt 基于int类型的自定义类型
type MyInt int

// NewInt int类型的别名
type NewInt = int

func main() {
	var a MyInt
	var b NewInt
	fmt.Printf("Type of a:%T Value:%v\n", a, a)
	fmt.Printf("Type of b:%T Value:%v\n", b, b)
}
