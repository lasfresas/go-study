package main

import "fmt"

func main() {
	var m1 map[string]int
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["小明"] = 18
	m1["白展堂"] = 25
	m1["姬无命"] = 26

	// delete删除元素
	delete(m1, "姬无命")
	fmt.Println(m1)
}
