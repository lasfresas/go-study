package main

import "fmt"

// 结构体的嵌套

// 用type关键字造的,都是类型
type Address struct {
	Province string
	City     string
}

type Person struct {
	Name    string
	Gender  string
	Age     int
	Address Address // 嵌套另外一个结构体。前面的是Person里面的字段名,后面的是指Address类型
}

func main() {
	p1 := Person{
		Name:   "小草莓",
		Gender: "男",
		Age:    18,
		Address: Address{
			Province: "上海",
			City:     "浦东新区",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Address)
	fmt.Println(p1.Address.Province, p1.Address.City)
}
