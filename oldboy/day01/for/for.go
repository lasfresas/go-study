package main

import "fmt"

// for循环
func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for range循环
	s := "hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}
}
