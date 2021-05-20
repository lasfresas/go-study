package main

import "fmt"

// 结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个person类型的变量p
	var p1 person
	// 通过字段赋值
	p1.name = "李想"
	p1.age = 18
	p1.gender = "男"
	p1.hobby = []string{"唱", "跳", "rap"}
	fmt.Println(p1)
	// 访问变量p的字段
	fmt.Printf("%T\n", p1) // main.person类型
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.gender)
	fmt.Println(p1.hobby)
}
