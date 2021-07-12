package main

import (
	"fmt"
	"io"
	"os"
)

// read by file.Read()
func readFile() {
	// 打开文件
	fileObj, err := os.Open("main.go") // 相对路径
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 提前注册，关闭文件
	defer fileObj.Close()
	// 使用file.Read()读取文件
	for {
		var tmp = make([]byte, 128)
		n, err := fileObj.Read(tmp)
		if err == io.EOF { // End Of File
			return
		}
		if err != nil {
			fmt.Printf("read by file.Read() failed, err:%v\n", err)
			return
		}
		fmt.Print(string(tmp[:n]))
	}
}

func main() {
	readFile()
}
