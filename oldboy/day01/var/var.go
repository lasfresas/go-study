package main

import "fmt"

// 全局变量m
var m = 100

func main() {
	var s1 = "zhang"
	fmt.Println(s1)
	// 类型推导，根据值判断变量的类型
	var s2 = "20"
	fmt.Println(s2)
	// 短变量声明，只能在函数中使用(局部变量)
	s3 := "heiheihei"
	fmt.Println(s3)
}
