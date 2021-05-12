package main

import "fmt"

// 流程控制之跳出与跳过for循环

func main() {
	// 当i=5时跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 跳出循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时跳过此次for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // 跳过本次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
