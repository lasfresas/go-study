package main

import (
	"encoding/json"
	"fmt"
)

// 反序列化：把json格式的字符串 --> Go能够识别的结构体
type point struct {
	// Go中如果标识符首字母是大写的，就表示对外可见（暴露的，公有的）
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	var p point
	str := `{"x":100,"y":200}`
	err := json.Unmarshal([]byte(str), &p)
	// 如果出错了
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	fmt.Print(p)
}
