package main

import (
	"fmt"
	"io"
	"os"
)

// 读文件1
func main() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed!, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
