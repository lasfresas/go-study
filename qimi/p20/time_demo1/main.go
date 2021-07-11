package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// 时间戳：1970.1.1 到现在的一个秒数
	timeStamp := now.Unix()
	fmt.Println(timeStamp)

	// 将时间戳转换为具体的时间格式
	t := time.Unix(timeStamp+3600, 0)
	fmt.Println(t)
}
