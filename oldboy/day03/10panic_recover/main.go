package main

import "fmt"

// panic和recover
/*
注意：
1. recover()必须搭配defer使用。
2. defer一定要在可能引发panic的语句之前定义。
*/

func funcA() {
	fmt.Println("a")
}

func funcB() {
	defer func() {
		err := recover()
		// 如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("数据库连接出错，释放连接……")
		}
	}()
	panic("数据库连接出错，程序退出") // 崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
