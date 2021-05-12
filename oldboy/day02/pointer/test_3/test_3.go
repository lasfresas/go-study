package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc000012088
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc000012088 type:*int
	fmt.Println(&b)                    // 0xc000006028
	fmt.Printf("type(&b):%T\n", &b)    // type(&b):**int
}
