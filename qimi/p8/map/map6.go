package main

import "fmt"

// map(映射)

func main() {
	// 值为切片的map
	var sliceMap = make(map[string][]int, 8) // 只完成了外层map的初始化
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		// sliceMap中没有中国这个键
		sliceMap["中国"] = make([]int, 8, 8) // 完成了对内层切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][3] = 300
	}
	// 遍历sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
