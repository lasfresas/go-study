package main

import "fmt"

// 二维数组
func main() {
	cityArray := [][]string{
		{"北京", "天津"},
		{"上海", "苏州"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(cityArray)       // 打印这个二维数组
	fmt.Println(cityArray[1][0]) // 拿出这个索引对应的元素
	fmt.Println()                // 空一行

	// 遍历二维数组
	for _, v1 := range cityArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
