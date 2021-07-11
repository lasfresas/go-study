package main

import "fmt"

func main() {
	var ch1 chan int        // 引用类型, 需要初始化之后使用
	ch1 = make(chan int, 1) // 带缓冲区通道==异步通道，无缓冲区通道==同步通道
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)
	close(ch1)
}
