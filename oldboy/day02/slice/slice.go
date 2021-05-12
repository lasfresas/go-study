package main

import "fmt"

// 切片

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	// 长度和容量
	/*
		切片指向了一个底层的数组
		切片的长度就是它元素的个数
		切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	*/

	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2.由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // 左闭右开
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	// 切片是引用类型, 指向了底层的一个数组
	a1[5] = 1111 // 修改了底层数组的值
	fmt.Println(s4)
}
