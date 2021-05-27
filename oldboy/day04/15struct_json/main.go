package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1. 序列化：把Go中的结构体变量 --> json格式的字符串
// 2. 反序列化：把json格式的字符串 --> Go能够识别的结构体

type person struct {
	// Go中如果标识符首字母是大写的，就表示对外可见（暴露的，公有的）
	Name string `json:"name" db:"name"` // 在json中叫name，在db数据库中叫name。键值对之间空格隔开
	Age  int    `json:"age"`
}

func main() {
	// 序列化
	p1 := person{
		Name: "小明",
		Age:  18,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) // 这里可以强转，因为字符串本来就是字符切片组成的

	// 反序列化
	var p2 person
	str := `{"name":"老八","age":30}`
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal()函数内部修改p2的值
	fmt.Printf("%v\n", p2)
}
