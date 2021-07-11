package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

// read by file.Read()
func readByFile() {
	// 打开文件
	fileObj, err := os.Open("./main.go") // 相对路径
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

// read by bufio
func readByBufio() {
	// 打开文件
	fileObj, err := os.Open("./main.go") // 相对路径
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

// read by ioutil
func readByIoutil() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read by ioutil failed, err:%v\n", err)
		return
	}
	fmt.Print(string(content)) // ioutil.ReadFile返回的是[]byte, 所以要转成string
}

// write
func write() {
	fileObj, err := os.OpenFile("./xxx.test", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	defer fileObj.Close() // 提前注册，关闭文件
	str := "大雕小哥哥\t"
	fileObj.Write([]byte(str)) // []byte
	fileObj.WriteString("日恁奶奶\n")
}

func main() {
	//readByFile()
	//readByBufio()
	//readByIoutil()
	write()
}
