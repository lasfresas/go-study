package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量的
	// 方式1:通过类型断言,一个个猜
	// 方法2:反射
	Obj := reflect.TypeOf(x)
	fmt.Println(Obj, Obj.Name(), Obj.Kind())
	//fmt.Printf("%T\n", Obj) // *reflect.rtype
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() // 拿到值对应的类型种类
	// 如何得到一个传入时候类型的变量
	switch k {
	case reflect.Int32:
		// 把反射取到的值转换成一个int32类型的变量
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
	case reflect.Float32:
		// 把反射取到的值转换成一个float32类型的变量
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem()用来获取指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(123)
	case reflect.Float32:
		v.Elem().SetFloat(123.456)
	}
}

type Cat struct{}

type Dog struct{}

func main() {
	// type of
	var a int8 = 1
	reflectType(a)
	var b float32 = 1.2
	reflectType(b)

	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)

	// value of
	var aa int32 = 12
	reflectValue(aa)
	var bb float32 = 12.34
	reflectValue(bb)

	// set value
	var aaa int32 = 100
	reflectSetValue(&aaa) // 传的是指针
	fmt.Println(aaa)
	var bbb float32 = 200
	reflectSetValue(&bbb)
	fmt.Println(bbb)
}
