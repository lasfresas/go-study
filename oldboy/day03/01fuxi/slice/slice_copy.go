package main

import "fmt"

// day03复习

// 切片是引用类型, 不存值
func main() {
	s1 := []int{1, 2, 3} // [1 2 3]
	s2 := s1             // 把s1的内存地址给了s2
	var s3 = make([]int, 3, 3)
	copy(s3, s1) // 把s1复制出一份副本, 拷贝给s3
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	s2[1] = 200     // s1和s2指向同一个内存地址, 故修改的是同一个
	fmt.Println(s1) // [1 200 3]
	fmt.Println(s2) // [1 200 3]
	fmt.Println(s3) // [1 2 3]
}
