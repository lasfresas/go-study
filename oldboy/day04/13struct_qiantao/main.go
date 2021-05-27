package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	addr address
}

type company struct {
	name string
	addr address
}

func main() {
	p1 := person{
		name: "小明",
		age:  9000,
		addr: address{
			province: "安徽",
			city:     "淮北",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)
	// fmt.Println(p1.city) //访问不了,需要匿名嵌套结构体
}
