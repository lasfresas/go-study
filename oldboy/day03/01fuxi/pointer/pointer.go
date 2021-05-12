package main

import "fmt"

// day03复习

// pointer
// Go里面的指针只能读不能修改, 不能修改指针变量指向的地址
func main() {
	addr := "沙河"
	addrP := &addr
	fmt.Println(addrP) // 内存地址
	fmt.Printf("%T\n", addrP)
	addrV := *addrP // 根据内存地址找值
	fmt.Println(addrV)
	fmt.Println(addr == addrV)
}
