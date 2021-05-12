package main

import "fmt"

// p6 流程控制 for
func main() {
	// 1.基本写法
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// 2.省略初始语句
	var j = 0
	for ; j < 10; j++ {
		fmt.Println(j)
	}

	fmt.Println()

	// 3.省略初始语句和结束语句
	var k = 10
	for k > 0 {
		fmt.Println(k)
		k--
	}
}
