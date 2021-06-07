package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("xxx.test", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello 小草莓\n") // 将数据先写入缓存
	}
	writer.Flush() // 将缓存中的内容写入文件
}
