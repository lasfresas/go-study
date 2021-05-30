package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段可见性 & 序列化和反序列化

// 如果一个Go语言包中定义的标识符是首字母大写的,那么就是对外可见的
type student struct {
	ID   int
	Name string
}

// student的构造函数
func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

// class(班级)的结构体
type class struct {
	Title    string    `json:"title"`
	Students []student `json:"students"`
}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "火箭101",
		Students: make([]student, 0, 20),
	}
	// 往班级c1中添加学生
	for i := 1; i <= 10; i++ {
		// 造10个学生
		tmpStu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("%#v\n", c1)

	// 序列化: Go语言中的数据 -> JSON格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed, err:", err)
		return
	}
	fmt.Printf("%T\n", data) // []uint8
	fmt.Printf("%s\n", data)

	// 反序列化: JSON格式的字符串 -> Go语言中的数据
	var c2 class
	jsonStr := `{"Title":"火箭101","Students":[{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"},{"ID":10,"Name":"stu10"}]}`
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed, err:", err)
		return
	}
	fmt.Printf("%#v\n", c2)
}
