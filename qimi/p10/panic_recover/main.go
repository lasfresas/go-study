package main

import "fmt"

// panic和recover

// 注意：
// 1. recover()必须搭配defer使用。
// 2. defer一定要在可能引发panic的语句之前定义。

func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func b error, the err is:", err)
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}
