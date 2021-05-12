package main

// package main表示这是个可执行文件，不是库

// 导入内置的fmt包
import "fmt"

// p2 基础
func main() {
	// 单行注释
	/*
		多行注释
	*/
	fmt.Print("牛年快乐！")
}

/*
1.go语言大括号不换行
2.语句结束不写分号
3.go install表示安装的意思，
它先编译源代码得到可执行文件，然后将可执行文件移动到GOPATH的bin目录下。
因为我们的环境变量中配置了GOPATH下的bin目录，所以我们就可以在任意地方直接执行可执行文件了。
*/
