package main

import (
	"fmt"
	"sort"
)

// 2.请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序（附加题，自行查资料解答）。
func main() {
	var a = [...]int{3, 7, 8, 9, 1}
	// a[:]得到的是一个切片，指向了底层的数组a
	sort.Ints(a[:])
	fmt.Println(a)
}
