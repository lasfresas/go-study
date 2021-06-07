package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello 小草莓"
	err := ioutil.WriteFile("./xxx.test", []byte(str), 0666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}
