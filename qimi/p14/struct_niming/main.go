package main

import "fmt"

// p14 匿名结构体
func main() {
	var user struct {
		name    string
		age     int
		married bool
	}
	user.name = "小草莓"
	user.age = 36
	user.married = true
	fmt.Println(user)
}
