package main

import (
	"fmt"
	"sync"
)

// 原生的map不是并发安全的

var (
	wg sync.WaitGroup
)

var m = make(map[int]int)

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			set(i, i+100)                              // map中设置键值对
			fmt.Printf("key:%v value:%v\n", i, get(i)) // 打印键值对
			wg.Done()
		}(i)
	}
	wg.Wait()
}
