package main

import (
	"fmt"
	"strings"
)

// Go中的string以原生数据类型出现，跟其他原生数据类型（int、float32、float64等）一样。
// Go中的string以UTF-8编码。
func main() {
	path := "Z:\\go-study\\oldboy\\day01\\string"
	fmt.Println(path)

	s1 := "I'm ok"
	fmt.Println(s1)
	// 多行的字符串
	s2 := `
世情薄
     人情恶
		  雨送黄昏花易落
`
	fmt.Print(s2)
	s3 := `Z:\go-study\oldboy\day01\string`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(path))
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(len(s3))
	// 拼接
	firstName := "zhang"
	secondName := "di"
	ss := firstName + secondName
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", firstName, secondName)
	fmt.Println(ss1)
	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "张三"))
}
