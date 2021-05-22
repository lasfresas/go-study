package main

import "fmt"

func main() {
	var a1 *int
	fmt.Println(a1) // nil, 没有分配内存地址
	var a2 = new(int)
	fmt.Println(a2)  // new是内置函数，返回一个指针
	fmt.Println(*a2) // a2这个地址所存的值，此时为int的零值
	*a2 = 100
	fmt.Println(*a2) // a2这个地址所存的值，此时为100
}
