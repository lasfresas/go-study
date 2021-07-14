package main

import (
	"fmt"
	"reflect"
)

// type of

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量的
	// 方式1：通过类型断言,一个个猜
	// 方法2：反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
	//fmt.Printf("%T\n", obj) // *reflect.rtype
}

type Cat struct{}

type Dog struct{}

func main() {
	var a int8 = 1
	reflectType(a) // int8
	var b float32 = 1.2
	reflectType(b) // float32
	var c Cat
	reflectType(c) // main.Cat
	var d Dog
	reflectType(d) // main.Dog
}
