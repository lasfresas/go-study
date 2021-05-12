package main

import "fmt"

// p5 运算符
func main() {
	// 1.算术运算符
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fmt.Println("***********************")

	// 2.关系运算符
	fmt.Println(10 > 2)  // true
	fmt.Println(10 != 2) // true
	fmt.Println(4 <= 5)  // true

	fmt.Println("***********************")

	// 3.逻辑运算符
	fmt.Println(10 > 5 && 1 == 1) // true
	fmt.Println(!(10 > 5))        // false
	fmt.Println(1 > 5 || 1 == 1)  // true

	fmt.Println("***********************")

	// 4.位运算符
	c := 1             // 001
	d := 5             // 101
	fmt.Println(c & d) // 按位与，001 => 1
	fmt.Println(c | d) // 按位或，101 => 5
	fmt.Println(c ^ d) // 异或，100 => 4

	fmt.Println("***********************")

	// 5.赋值运算符
}
