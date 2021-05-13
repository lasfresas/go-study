package main

import "fmt"

// defer面试题
// 提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

// 1. defer calc("1", 1, calc("10", 1, 2))
// 2. calc("10", 1, 2)  // "10" 1 2 3
// 3. a = 0
// 4. defer calc("2", 0, calc("20", 0, 2))
// 5. calc("20", 0, 2)  // "20" 0 2 2
// 6. b = 1
// 7. calc("2", 0, 2)  // "2" 0 2 2
// 8. calc("1", 1, 3)  // "1" 1 3 4

/*
	最终的结果：
	"10" 1 2 3
	"20" 0 2 2
	"2" 0 2 2
	"1" 1 3 4
*/
