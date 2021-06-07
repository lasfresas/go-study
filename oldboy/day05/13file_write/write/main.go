package main

// Write和WriteString

import (
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
	str := "你好\n"
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小草莓") //直接写入字符串数据
}
