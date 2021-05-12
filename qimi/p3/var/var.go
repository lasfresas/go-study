package main

import "fmt"

// 函数外的每个语句都必须以关键字开始（var、const、func等）

// p3 变量
func main() {
	// 标准声明格式
	var age int
	var name string
	var isOk bool
	fmt.Println(age, name, isOk)
	// 批量声明变量
	var (
		a string
		b int
		c bool
		d float32
	)
	fmt.Println(a, b, c, d)
	// 声明变量的同时指定初始值
	var name1 string = "小明"
	var age1 int = 18
	fmt.Println(name1, age1)
	// 类型推导，让编译器根据变量的初始值推导出其类型
	var name2 = "小红"
	var age2 = 19
	fmt.Println(name2, age2)
	// 短变量声明，只能用在函数中
	m := 10
	fmt.Println(m)
}
