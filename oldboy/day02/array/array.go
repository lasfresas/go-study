package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量（长度）
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 不初始化, 默认0值
	fmt.Println(a1, a2)
	// 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 初始化方式2
	// 根据初始值自动推断长度
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 初始化方式3
	// 根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)
}
