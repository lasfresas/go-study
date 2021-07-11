package main

import (
	"fmt"
	"sync"
)

// goroutine demo 3

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) { // 这个匿名函数还是个闭包(包含了外部的变量i)
			fmt.Println("hello,", i)
			wg.Done()
		}(i) // 传一个副本进来, 保证一致性
	}
	fmt.Println("hello, main")
	wg.Wait()
}
