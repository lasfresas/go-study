package main

import (
	"fmt"
	"io/ioutil"
)

// read by ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("main.go") // 相对路径
	if err != nil {
		fmt.Printf("read by ioutil failed, err:%v\n", err)
		return
	}
	fmt.Print(string(content)) // ioutil.ReadFile返回的是[]byte, 所以要转成string
}

func main() {
	readByIoutil()
}
