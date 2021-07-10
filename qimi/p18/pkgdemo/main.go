package main

import (
	"fmt"

	jisuan "pkgdemo/calc" // 可以给引入的包起个别名
)

func main() {
	fmt.Println("hello", jisuan.Name)
	fmt.Println(jisuan.Add(100, 200))
}

func init() {
	fmt.Println("main.init()")
}
