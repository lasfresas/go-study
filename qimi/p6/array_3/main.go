package main

import "fmt"

// 数组是值类型

func f1(a [3]int) {
	a[0] = 100
}

func main() {
	x := [3]int{1, 2, 3}
	fmt.Println(x)
	f1(x)
	fmt.Println(x)
}
