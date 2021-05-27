package main

// 结构体嵌套
type addr struct {
	province string
	city     string
}

type student struct {
	name    string
	address addr // 嵌套别的结构体
}
