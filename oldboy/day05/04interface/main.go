package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车, 都能跑

// 定义一个car接口类型
// 不管是什么结构体，只要有run方法都能是car类型
type car interface {
	run()
}

type ferrari struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s速度70迈\n", f.brand)
}

type porsche struct {
	brand string
}

func (p porsche) run() {
	fmt.Printf("%s速度700迈\n", p.brand)
}

// drive函数接收一个car类型的变量
func drive(c car) {
	c.run()
}

func main() {
	var f1 = ferrari{
		brand: "法拉利",
	}
	var p1 = porsche{
		brand: "保时捷",
	}

	drive(f1)
	drive(p1)
}
