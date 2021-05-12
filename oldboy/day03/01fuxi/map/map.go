package main

import "fmt"

// day03复习

// map
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["lixiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["lixiang"])
	fmt.Println(m1["jiwuming"]) // 如果key不存在, 返回的是value对应类型的0值
	// 如果返回值是bool型，我们通常用ok接收它
	score, ok := m1["jiwuming"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是", score)
	}

	delete(m1, "lixiang")
	fmt.Println(m1)
}
