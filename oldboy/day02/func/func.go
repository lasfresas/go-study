package main

import "fmt"

// 函数
// 是一段代码的封装
// 让代码结构更清晰简洁

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值
func f3() int {
	return 3
}

// 参数可以命名也可以不命名
func f4(x int, y int) (ret int) {
	ret = x + y
	return // 命名的返回值，可以只写return
}

// 多个返回值
func f5() (int, string) {
	return 1, "傻逼"
}

// Go语言中没有默认参数这个概念
func main() {
	r := f4(10, 20)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)
}
