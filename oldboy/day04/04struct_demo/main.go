package main

import "fmt"

// 结构体是值类型(复制、粘贴)
type person struct {
	name, gender string
}

// Go语言中函数传参永远是拷贝，函数内部永远是副本
func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女" // 根据内存地址找到原来的变量，修改的就是原来的变量
	x.gender = "女" // 语法糖，自动根据指针找对应的变量
}
func main() {
	var p person
	p.name = "xiaoming"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) // 男
	f2(&p)
	fmt.Println(p.gender) // 女

	// 结构体指针1
	var p2 = new(person) // 使用new函数得到的是一个类型的指针
	// (*p2).name = "理想"
	p2.name = "理想"          // 语法糖
	fmt.Printf("%T\n", p2)  // *main.person
	fmt.Printf("%p\n", p2)  // p2保存的值
	fmt.Printf("%p\n", &p2) // p2自己的内存地址

	// 结构体指针2
	// 1. key-value初始化
	var p3 = &person{
		name:   "小阿giao",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 2. 使用值列表的形式初始化（值的顺序跟结构体定义的字段顺序相同）
	p4 := &person{
		"郭老师",
		"女",
	}
	fmt.Printf("%#v\n", p4)
}
