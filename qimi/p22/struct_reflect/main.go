package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
/*
反射是把双刃剑:
	1. 基于反射的代码是很脆弱的, 反射中的错误会在真正运行的时候才会引发panic
	2. 大量使用反射的代码通常难以理解
	3. 反射的性能低下
*/

type student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

// 根据反射去获取结构体中方法的函数
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println(t.NumMethod()) // 取到变量的方法数量

	v := reflect.ValueOf(x)
	// 因为下面需要拿到具体的方法去调用, 所以使用的是值信息
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args) // 调用方法
	}
}

func main() {
	stu1 := student{
		Name:  "小草莓",
		Score: 90,
	}

	// 通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())

	// 遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		// 根据结构体字段的索引去取字段
		fieldObj := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println(fieldObj.Tag.Get("json"), fieldObj.Tag.Get("ini"))
	}

	// 根据名字去取结构体中的字段
	field, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v\n", field.Name, field.Type, field.Tag)
	}

	printMethod(stu1)
}
