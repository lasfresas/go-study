package main

import "fmt"

func main() {
	var a *int      // nil, 未分配的的内存地址
	*a = 100        // 对一个未分配的内存地址赋值为100
	fmt.Println(*a) // panic: runtime error: invalid memory address or nil pointer dereference
}
