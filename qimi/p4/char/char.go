package main

import "fmt"

// p4 字符
func main() {
	// byte	uint8的别名
	// rune int32的别名
	var c1 byte = 'c'
	var c2 rune = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T  c2:%T\n", c1, c2)

	s := "Google"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\t", s[i])
	}
}
