package main

import (
	"encoding/json"
	"fmt"
)

// 序列化：把Go中的结构体变量 --> json格式的字符串
type point struct {
	// Go中如果标识符首字母是大写的，就表示对外可见（暴露的，公有的）
	X int `json:"x"`
	Y int `json:"y"`
}

func main() {
	p := point{100, 200}
	// Go语言中把错误当成值返回,错误通常作为第二个返回值
	str, err := json.Marshal(p)
	// 如果出错了
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Print(string(str))
}
