package main

import (
	"fmt"
)

// map(映射)
// map是无序的
func main() {
	// 声明但不初始化,map的零值就是nil
	var a map[string]int
	fmt.Println(a == nil)
	// map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// map中添加键值对
	a["小明"] = 100
	a["小强"] = 200
	fmt.Println(a)
	fmt.Printf("Type:%T\n", a) // Type:map[string]int
	// 声明map的同时,完成初始化
	b := map[int]bool{
		1: true,
		2: false,
		3: true,
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("Type:%T\n", b) // Type:map[int]bool
}
