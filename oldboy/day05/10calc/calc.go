package calc

import "fmt"

/*
init()函数：
	每个包导入的时候自动执行，它没有参数没有返回值也不能手动调用
init()函数的执行时机：
	全局声明 -> init() -> main()
*/
func init() {
	fmt.Println("import时自动执行...")
}

// 包中的标识符大写才能公有(可以被外部的包调用), 小写就是私有(只能在当前的包中使用)
func Add(x, y int) int {
	return x + y
}
