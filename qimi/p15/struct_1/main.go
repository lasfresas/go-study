package main

import "fmt"

// 结构体的匿名字段

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"小草莓",
		20,
	}
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}
