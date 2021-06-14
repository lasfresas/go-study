package main

import "fmt"

// goroutine channel

/*
两个goroutine，两个channel
	1. 生成0-100的数发送到ch1
	2. 从ch1中取出数据计算它的平方，把结果发送到ch2中
*/

// 1. 生成0-100的数发送到ch1
func f1(ch1 chan<- int) { // 参数里面的符号表示数据只能往里面入
	for i := 0; i <= 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// 2. 从ch1中取出数据计算它的平方，把结果发送到ch2中
func f2(ch1 <-chan int, ch2 chan<- int) { // 参数里面的符号表示限制数据只能从ch1里面出，数据只能往ch2里面入
	// 从通道中取值的方式1
	for {
		tmp, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- tmp * tmp
	}
	close(ch2)
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	// 从通道中取值的方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
