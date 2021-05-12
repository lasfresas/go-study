package main

import "fmt"

// day03复习

// 切片slice
// 切片是引用类型, 本质是对底层数组的封装
func main() {
	var s1 []int           // 没有分配内存==nil
	fmt.Println(s1)        // []
	fmt.Println(s1 == nil) // true
	s1 = []int{1, 2, 3}    // 直接传值, 初始化
	fmt.Println(s1)

	// make初始化, 分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2) // [false false]

	s3 := make([]int, 0, 4)
	fmt.Println(s3 == nil) // false, 因为s3已经用make初始化了
}
