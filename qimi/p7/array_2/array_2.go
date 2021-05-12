package main

import "fmt"

// p7 二维数组
func main() {
	cityArray := [][]string{
		{"北京", "西安"},
		{"上海", "杭州"},
		{"成都", "重庆"},
		{"广州", "深圳"},
	}
	fmt.Println(cityArray)       // 打印这个二维数组
	fmt.Println(cityArray[1][0]) // 拿出这个索引对应的元素
	fmt.Println()                // 空一行

	// 二维数组的遍历
	for _, v1 := range cityArray {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
