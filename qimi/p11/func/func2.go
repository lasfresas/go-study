package main

import "fmt"

// p11 闭包

// 定义一个函数，它的返回值是函数
// 把函数作为返回值
func a() func() {
	name := "沙河娜扎"
	return func() {
		fmt.Println("hello", name)
	}
}

func main() {
	// 闭包 = 函数 + 外层变量的引用
	r := a() // r此时就是一个闭包
	r()      // 相当于执行了a函数内部的匿名函数
}
