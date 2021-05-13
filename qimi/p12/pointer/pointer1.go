package main

import "fmt"

// p12 指针
// &(取地址)、*(根据地址取值)
// 区别于C/C++中的指针，Go语言中的指针不能进行偏移和运算，是安全指针。

func main() {
	a := 10
	b := &a
	fmt.Printf("%v %p\n", a, &a)
	fmt.Printf("%v %p\n", b, b)
	// 变量b是一个int类型的指针(*int)，它存储的是变量a的内存地址
	c := *b // 根据指针地址取值
	fmt.Println(c)
}
