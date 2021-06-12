package main

import (
	"fmt"
	"sync"
)

// goroutine demo 2

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() { // 这个匿名函数还是个闭包(包含了外部的变量i)
			fmt.Println("hello,", i)
			wg.Done()
		}()
	}
	fmt.Println("hello, main")
	wg.Wait()
}
