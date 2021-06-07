package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读文件2
// bufio按行读取
func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 注意是字符
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}
