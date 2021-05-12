package main

import (
	"fmt"
	"unicode"
)

// 1.判断字符串中汉字的个数
func main() {
	// 难点是判断一个字符是汉字
	s1 := "汉语にほんご조선어English"
	// 1.依次拿到字符串中的字符
	var count int
	for _, c := range s1 {
		// 2.判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	// 3.把汉字出现的次数累加得到最终结果
	fmt.Println(count)
}
