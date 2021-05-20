package main

import "fmt"

// 匿名结构体：多用于临时场景
var s struct {
	x string
	y int
}

func main() {
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}
