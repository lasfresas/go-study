package main

import "fmt"

// map和slice组合
func main() {
	// map类型元素的切片（切片的元素是map）
	var s1 = make([]map[int]string, 10, 10) // 初始化了切片，但没有初始化里面的map类型元素
	s1[0] = make(map[int]string, 1)         // 初始化切片s1的第0个元素(map类型)，map的容量为1
	s1[0][10] = "沙河"                        // 切片s1的第0个元素(map类型)，key为10赋值为"沙河"
	fmt.Println(s1)

	// 值为切片类型的map（map的value是切片）
	var m1 = make(map[string][]int, 0)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
