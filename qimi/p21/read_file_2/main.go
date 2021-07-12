package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// read by bufio
func readByBufio() {
	// 打开文件
	fileObj, err := os.Open("main.go") // 相对路径
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 提前注册，关闭文件
	defer fileObj.Close()
	// 使用bufio来读取文件
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}

func main() {
	readByBufio()
}
