package main

import (
	"fmt"
	"time"
)

/*
	时间类型有一个自带的方法Format进行格式化
	需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。
	也许这就是技术人员的浪漫吧。
*/
func main() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分05秒
	ret := now.Format("2006-01-02 15:04:05")
	fmt.Println(ret)
}
