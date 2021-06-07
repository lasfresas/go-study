package main

import (
	"bufio"
	"fmt"
	"os"
)

func useBufio() {
	fmt.Print("请输入:")
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Printf("你输入的内容是:%s\n", s)
}

func main() {
	useBufio()
}
