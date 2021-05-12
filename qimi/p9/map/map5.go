package main

import "fmt"

// p9 map

func main() {
	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) // 只是完成了切片的初始化
	fmt.Println(mapSlice)
	fmt.Println(mapSlice[0] == nil)

	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["小米"] = 3999
	fmt.Println(mapSlice)
}
