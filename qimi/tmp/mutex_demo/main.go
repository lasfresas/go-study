package main

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 释放锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
