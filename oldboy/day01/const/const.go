package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会修改
const pi = 3.1415926

// 批量声明常量
const (
	ok       = 200
	notfound = 404
)

// 批量声明常量时，如果某一行声明后没有赋值，默认和上一行一样
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println(n1, n2, n3)
}
