package main

import "fmt"

// map(映射)

func main() {
	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) // 只完成了外层切片的初始化
	fmt.Println(mapSlice[0] == nil)
	fmt.Println(mapSlice)

	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["小米"] = 3999
	fmt.Println(mapSlice)
}
