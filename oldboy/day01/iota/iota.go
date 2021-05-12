package main

import "fmt"

// iota，类似枚举
/*
iota在const关键字出现时重置为0。const中每新增一行常量的声明将使iota计数1次
使用iota能简化定义，在定义枚举的时候很有用
*/
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// 插队
const (
	b1 = iota // 0
	b2 = 100  //100
	b3 = iota // 2
	b4        // 3
)

// 多个变量声明在一行
const (
	c1, c2 = iota + 1, iota + 2 // c1=1, c2=2
	c3, c4 = iota + 1, iota + 2 // c3=2, c4=3
)

// 定义数量级
const (
	_  = iota             // iota=0, 扔了
	KB = 1 << (10 * iota) // iota=1, KB=2^10B
	MB = 1 << (10 * iota) // iota=2, MB=2^20B
	GB = 1 << (10 * iota) // iota=3, GB=2^30B
	TB = 1 << (10 * iota) // iota=4, TB=2^40B
	PB = 1 << (10 * iota) // iota=5, PB=2^50B
)

func main() {
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println()
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b4:", b4)
	fmt.Println()
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
}
