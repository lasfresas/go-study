package main

import "fmt"

// p4 字符串常见操作
func main() {
	// 求字符串的长度，len()方法
	s1 := "hello"
	fmt.Println(len(s1))
	// 汉字，每个字占3字节
	s2 := "中国"
	fmt.Println(len(s2))

	// 字符串拼接
	fmt.Println(s1 + s2)
}
