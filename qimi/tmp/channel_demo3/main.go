package main

import "fmt"

// worker pool（goroutine池）
/*
在工作中我们通常会使用可以指定启动的goroutine数量 —— worker pool模式，控制goroutine的数量，防止goroutine泄漏和暴涨。
*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, job)
		results <- job * 2
		fmt.Printf("worker:%d stop job:%d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}
	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
