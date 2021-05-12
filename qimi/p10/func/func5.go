package main

import "fmt"

// p10 函数可以作为返回值

func sum(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	r1 := calc(100, 200, sum)
	fmt.Println(r1)
	r2 := calc(100, 200, sub)
	fmt.Println(r2)
}
