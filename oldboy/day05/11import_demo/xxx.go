package main

import (
	// 前面的calc是给包起的别名
	calc "../10calc"
	"fmt"
	// 前面加下划线：导入，但不用。诶，就是玩(导入包，而不使用包内部的数据)
	_ "os"
)

func init() {
	fmt.Println("自动执行！")
}

func main() {
	ret := calc.Add(6, 4)
	fmt.Println(ret)
}
