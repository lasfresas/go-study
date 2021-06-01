package main

import "fmt"

/*
	想要判断空接口中的值这个时候就可以使用类型断言，其语法格式：
	x.(T)
	x：表示类型为interface{}的变量
	T：表示断言x可能是的类型。
	该语法返回两个参数，第一个参数是x转化为T类型后的变量，第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败。
*/

func main() {
	var x interface{}
	x = "Hello 沙河"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
