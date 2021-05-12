package main

import "fmt"

// p10 函数进阶之变量作用域

// 定义全局变量num
var num int = 10

// 定义函数
func testGlobal() {
	num := 100
	// 可以在函数中访问全局变量
	// 1. 先在自己函数中查找，找到了就用
	// 2. 找不到就往外找全局变量
	fmt.Println("变量num", num)
}

func main() {
	testGlobal()
}
