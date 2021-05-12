package main

import "fmt"

// p10 函数

// 定义一个不需要参数也没有返回值的函数：sayHello
func sayHello() {
	fmt.Println("Hello,沙河小王子！")
}

// 定义一个接收string类型的name参数
func sayHello2(name string) {
	fmt.Println("Hello", name)
}

// 定义接收多个参数的函数并且有一个返回值
func intSum(a int, b int) int {
	ret := a + b
	return ret
}
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return
}

// 函数接收可变参数,在参数名后面加... 表示可变参数
// 可变参数在函数体中是切片类型
func intSum3(a ...int) int {
	ret := 0
	for _, arg := range a {
		ret = ret + arg
	}
	return ret
}

func main() {
	// 函数调用
	r1 := intSum3()
	r2 := intSum3(10)
	r3 := intSum3(10, 20)
	r4 := intSum3(10, 20, 30)
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)
}
