package main

import "fmt"

// p14 结构体初识

// 结构体本质上是一种聚合型的数据类型
type person struct {
	name string
	age  int
	city string
}

func main() {
	var p person
	p.name = "小草莓"
	p.age = 26
	p.city = "上海"
	fmt.Printf("姓名:%s\n年龄:%d\n城市:%s\n", p.name, p.age, p.city)
	fmt.Println(p)
}
