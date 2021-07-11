package main

import (
	"fmt"
	"sync"
)

// goroutine demo
var wg sync.WaitGroup

func hello(i int) {
	wg.Add(1) // 计数器+1
	fmt.Println("hello,", i)
	wg.Done() // 计数器-1
}

func main() { // 开启一个主goroutine去执行main函数
	for i := 0; i < 10; i++ {
		go hello(i) // 开启了一个goroutine去执行hello这个函数
	}
	fmt.Println("hello, main!")
	wg.Wait() // 阻塞, 等所有小弟都干完活再收兵
}
