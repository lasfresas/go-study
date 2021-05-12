package main

import "fmt"

// day03复习
func main() {
	var name string
	name = "理想"
	fmt.Println(name)
	var ages [30]int // 声明了一个变量ages, 它是[30]int类型
	ages = [30]int{1, 2, 3, 4, 5, 6}
	fmt.Println(ages)
	var ages2 = [...]int{1, 2, 3, 4}
	fmt.Println(ages2)
	var ages3 = [...]int{1: 100, 99: 200}
	fmt.Println(ages3)

	// 二维数组
	var a1 [3][2]int // [[0 0] [0 0] [0 0]]
	a1 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)

	// 数组是值类型
	x := [3]int{1, 2, 3}
	fmt.Println(x) // [1 2 3]
	f1(x)
	fmt.Println(x) // [1 2 3]
}

func f1(a [3]int) {
	// Go语言中的函数传递的都是值(Ctrl+C Ctrl+V)
	a[1] = 100 // 此处修改的是副本的值
}
