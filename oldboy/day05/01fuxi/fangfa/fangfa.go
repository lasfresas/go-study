package main

import "fmt"

type person struct {
	name string
	age  int
}

// 方法是有接收者的函数，接收者指的是哪个类型的变量可以调用这个函数
// 接收者使用对应类型的首字母小写
// 指定了接收者之后，只有接收者这个变量可以调用方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.name, str)
}

func (p person) false_guonian() {
	p.age++ // 此处改的是副本
}

/*
指针接收者
	1. 需要修改结构体变量的值时
	2. 结构体本身比较大，拷贝的内存开销大
	3. 保持一致性：如果有一个方法使用了指针接收者，其他的方法为了统一也要使用
*/
func (p *person) true_guonian() {
	p.age++ // 此处改的是对应内存地址的原变量
}

func main() {
	var p1 person
	p1.name = "小明"
	p1.age = 20

	p2 := person{"小红", 30}

	p1.dream("做个咸鱼")
	p2.dream("赚大钱")
	fmt.Println(p1.age)
	p1.false_guonian() // 假过年
	fmt.Println(p1.age)
	p1.true_guonian() // 真过年
	fmt.Println(p1.age)
}
