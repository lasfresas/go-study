package main

import "fmt"

// map(映射)

func main() {
	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["小米"] = 3999
	scoreMap["苹果"] = 5999

	// 判断"华为"在不在scoreMap中
	v, ok := scoreMap["华为"] // 如果map中存在这个key：那么ok返回true，value就赋值给v；否则ok返回false，value的零值赋值给v。
	if ok {
		fmt.Println("增智慧", v)
	} else {
		fmt.Println("连哄带蒙", v)
	}
}
