package main

import "fmt"

// 递归:函数自己调用自己
// 递归适合处理:相同、问题规模越来越小的场景

// 计算n的阶乘
func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(7)
	fmt.Println(ret)
}
