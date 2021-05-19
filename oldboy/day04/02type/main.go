package main

import "fmt"

// 自定义类型和类型别名
type myInt int    // 自定义类型
type newInt = int // 类型别名

func main() {
	var m myInt
	m = 100
	var n newInt
	n = 100
	fmt.Printf("%T\n", m)
	fmt.Printf("%T\n", n)
}
