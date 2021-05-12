package main

import (
	"fmt"
	"strings"
)

// 2. how do you do 单词出现的次数
func main() {
	s1 := "how do you do"
	// 2.1 把字符串按照空格切割，得到切片
	s2 := strings.Split(s1, " ")
	// 2.2 遍历切片存储到一个map
	m1 := make(map[string]int, 10)
	for _, w := range s2 {
		if _, ok := m1[w]; !ok {
			// 1. 如果原来map中不存在w这个key, 那么出现次数=1
			m1[w] = 1
		} else {
			// 2. 如果map中存在w这个key, 那么出现次数+1
			m1[w]++
		}
	}
	// 2.3 累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}
