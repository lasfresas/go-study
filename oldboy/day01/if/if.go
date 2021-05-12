package main

import "fmt"

// if条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("该写作业了！")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("保温杯里泡枸杞")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习")
	}
}
