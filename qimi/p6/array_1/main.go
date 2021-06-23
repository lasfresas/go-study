package main

import "fmt"

// 数组
func main() {
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("*********1.数组的初始化**********")
	// 1. 数组的初始化
	var cityArray = [4]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(cityArray)

	fmt.Println("*********2.编译器推导数组长度**********")
	// 2. 编译器推导数组长度
	var boolArray = [...]bool{true, false, false, true, false}
	fmt.Println(boolArray)

	fmt.Println("*********3.使用索引值方式初始化**********")
	// 3. 使用索引值方式初始化
	// 这里Java是数组的第2个元素，因为它索引值为1。前面会有一个空的字符串元素
	var langArray = [...]string{1: "Java", 2: "Go", 3: "Python", 6: "Ruby"}
	fmt.Println(langArray)
	// 上面的数组长度为7，类型为“[7]string”
	fmt.Printf("%T\n", langArray)

	fmt.Println("*******for遍历********")
	// 普通遍历
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}

	fmt.Println("*******for range遍历********")
	// for range遍历
	for _, value := range cityArray {
		fmt.Println(value)
	}
}
