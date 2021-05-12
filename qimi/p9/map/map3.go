package main

import "fmt"

// p9 map

func main() {
	// 遍历map
	var scoreMap = make(map[string]int, 8)
	scoreMap["红米"] = 1999
	scoreMap["小米"] = 3999
	scoreMap["苹果"] = 5999

	// for range
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	// 只遍历key
	for k := range scoreMap {
		fmt.Println(k)
	}
	// 只遍历value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// delete()删除键值对
	delete(scoreMap, "红米")
	fmt.Println(scoreMap)
}
