package main

import (
	"fmt"
	"os"
)

// write
func write() {
	fileObj, err := os.OpenFile("xxx.test", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close() // 提前注册，关闭文件
	str := "大雕小哥哥\t"
	fileObj.Write([]byte(str))    // []byte
	fileObj.WriteString("日恁奶奶\n") // string
}

func main() {
	write()
}
