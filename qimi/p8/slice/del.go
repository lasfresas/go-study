package main

import "fmt"

// p8 切片删除元素
func main() {

	a := []string{"北京", "上海", "广州", "深圳"}
	// 删除"广州"，索引为2，故先切片0:2，注意是左闭右开区间
	a = append(a[0:2], a[3])
	fmt.Println(a)
}
