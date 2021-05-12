package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// p9 map

func main() {
	// 按照某个固定顺序遍历map
	var scoreMap = make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99的随机整数
		scoreMap[key] = value
	}

	// 按照key从小到大的顺序去遍历scoreMap
	// 1. 先取出所有的key存放到切片中
	keys := make([]string, 0, 100)
	for k := range scoreMap {
		keys = append(keys, k)
	}
	// 2. 对key做排序
	sort.Strings(keys) // keys目前是一个有序的切片
	// 3. 按照排序后的key对scoreMap排序
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	// but，为什么随机生成了固定的数？甚至和视频里一样？
}
