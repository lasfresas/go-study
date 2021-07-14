package main

import (
	"fmt"
	"reflect"
)

// set value

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	switch v.Elem().Kind() { // Elem()用来获取指针取对应的值
	case reflect.Int32:
		v.Elem().SetInt(123)
	case reflect.Float32:
		v.Elem().SetFloat(123.123)
	}
}
func main() {
	var aaa int32 = 1
	reflectSetValue(&aaa) // 传的是指针
	fmt.Println(aaa)
	var bbb float32 = 1
	reflectSetValue(&bbb)
	fmt.Println(bbb)
}
