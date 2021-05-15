package main

import "fmt"

// fmt

func main() {
	var m1 = make(map[string]int, 1)
	m1["理想"] = 100
	// %v	值的默认格式表示
	// %+v	类似%v，但输出结构体时会添加字段名
	// %#v	值的Go语法表示
	fmt.Printf("%v\n", m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%#v\n", m1)
}
