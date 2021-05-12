package main

import "fmt"

// byte和rune类型
// Go中为了处理非ASCII码类型的字符，定义了新的rune类型(int32)

func main() {
	c1 := 'R'
	c2 := "R"
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := '红'
	c4 := "红"
	fmt.Printf("c3:%T c4:%T\n", c3, c4)

	// 类型转换
	n := 10 // int
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
