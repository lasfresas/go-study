package main

import "fmt"

// 在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
// 第一步：返回值赋值
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
	// 1. 返回值赋值为5
	// 2. x++变成6
	// 3. RET返回值5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
	// 1. 返回值x赋值为5
	// 2. 返回值x++变成6
	// 3. RET返回值6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
	// 1. y=x赋值为5
	// 2. x++变成6
	// 3. RET返回值5(y)
}

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return 5 // 返回值=x=5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
