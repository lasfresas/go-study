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
	// 2. 执行defer语句，x++变成6
	// 3. RET返回值5
	// Go的函数传参是值拷贝。返回值是x的一个副本，和x本身没关系了
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
	// 1. 返回值x赋值为5
	// 2. 执行defer语句，返回值x++变成6
	// 3. RET返回值6
	// 这和f1()有区别，因为这里f2()在定义的时候指明了返回值为x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
	// 1. x拷贝一份y，赋值为5
	// 2. 执行defer语句，x++变成6，y为5不变
	// 3. RET返回值y，即为5
	// 指明了返回值为y
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x) // Go的函数传参是值拷贝。外面传进去x，外面的x在里面复制出一个x副本。改变的是函数中的副本。
	return 5 // 返回值=x=5
}

// 传一个x的指针到匿名函数中
func f5() (x int) {
	defer func(x *int) {
		(*x)++
	}(&x)
	return 5 // 1. 返回值x=5  2. defer x=6  3. RET返回
}

func main() {
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 6
}
