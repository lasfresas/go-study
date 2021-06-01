package main

import "fmt"

/*
空接口是指没有定义任何方法的接口。
因此任何类型都实现了空接口。
空接口类型的变量可以存储任意类型的变量。
*/

// 空接口的应用：
// 1. 空接口类型作为函数的参数
// 2. 空接口类型可以作为map的value

func main() {
	// 定义一个空接口
	var x interface{}
	s := "hello, 小草莓"
	x = s
	fmt.Printf("type:%T\tvalue:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T\tvalue:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T\tvalue:%v\n", x, x)

	var m = make(map[string]interface{}, 16)
	m["name"] = "娜扎"
	m["age"] = 18
	m["hobby"] = []string{"唱", "跳", "rap", "篮球"}
	fmt.Println(m)
}
