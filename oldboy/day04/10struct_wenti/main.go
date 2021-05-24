package main

import "fmt"

// 结构体遇到的问题

// 结构体是值类型，赋值的时候都是拷贝。
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

/*
构造函数：约定成俗用new开头
返回的是结构体还是结构体指针
当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
*/

// 为什么要有构造函数？
// 调用该构造函数，可以返回一个person类型的变量
func newPerson(n string, a int) *person {
	tmp := &person{
		name: n,
		age:  a,
	}
	return tmp
}

// 调用该构造函数，可以返回一个dog类型的变量
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("小明", 18)
	p2 := newPerson("小红", 19)
	fmt.Println(p1, p2)
	d1 := newDog("皮皮")
	fmt.Println(d1)
}
