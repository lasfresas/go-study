package main

import "fmt"

// p9 map

func main() {
	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["小米"] = 3999
	scoreMap["苹果"] = 5999

	// 判断"华为"在不在scoreMap中
	value, ok := scoreMap["华为"]
	if ok {
		fmt.Println("华为为我增智慧", value)
	} else {
		fmt.Println("哄蒙哄蒙，连哄带蒙", value)
	}
}
