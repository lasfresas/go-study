package main

// 可以给引入的包起个别名
import (
	jisuan "../calc"
	"fmt"
)

func main() {
	fmt.Println("hello", jisuan.Name)
	ret := jisuan.Add(10, 20)
	fmt.Println(ret)
}
