package main

import "fmt"

func assign(a interface{}) {
	str, ok := a.(string) // 类型断言
	if ok {
		fmt.Println("传进来的是一个字符串:", str)
	} else {
		fmt.Println("猜错了")
	}
}

func main() {
	assign("100")
}
