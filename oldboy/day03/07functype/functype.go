package main

import "fmt"

// 函数也是一种类型, 它可以作为参数, 也可以作为返回值

func f1() {
	fmt.Println("Hello 沙河")
}
func f2() int {
	return 10
}

// 函数可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数还可以作为返回值
func f5(x func() int) func(int, int) int {
	return ff
}
func ff(a, b int) int {
	return a + b
}

func main() {
	a := f1
	fmt.Printf("%T\n", a) // func()
	b := f2
	fmt.Printf("%T\n", b) // func() int

	f3(f2) // 10
	f3(b)  // 10

	fmt.Printf("%T\n", f4) // func(int, int) int

	f7 := f5(f2)
	fmt.Printf("%T\n", f7) // func(int, int) int
}
