package main

import "fmt"

// 3.回文判断
func main() {
	ss := "上海自来水来自海上"
	// 解题思路
	// 把字符串中的字符拿出来放到一个[]rune切片中
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println("[]rune:", r)

	for i := 0; i < len(r)/2; i++ {
		// 上 r[0]  r[len(r)-1]
		// 海 r[1]  r[len(r)-1-1]
		// 自 r[2]  r[len(r)-1-2]
		// 来 r[3]  r[len(r)-1-3]
		// ...
		// c r[i]  r[len(r)-1-i]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
		}
	}
	fmt.Println("是回文")
}
