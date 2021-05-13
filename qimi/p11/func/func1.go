package main

import "fmt"

// p11 匿名函数

func main() {
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()
}
