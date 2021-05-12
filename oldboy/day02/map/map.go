package main

import "fmt"

// map类型的变量默认初始值为nil，需要使用make()函数来分配内存。
// make(map[KeyType]ValueType, [cap])
// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。
func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化（没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["小明"] = 18
	m1["白展堂"] = 25
	m1["姬无命"] = 26

	fmt.Println(m1)
	fmt.Println(m1["小明"])
	value, ok := m1["小米"]
	// 约定成俗用ok接收返回的bool值
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
}
