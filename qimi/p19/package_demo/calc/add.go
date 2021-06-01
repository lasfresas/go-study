package calc

import (
	"../snow"
	"fmt"
)

// 标识符首字母大写表示对外可见

// Name 定义一个测试的全局变量
var Name = "小草莓"

// Add 是一个计算两个int类型数据的和的函数
func Add(x, y int) int {
	snow.Snow()
	Sub(x, y) // 同一个包中的不同文件可以直接调用
	return x + y
}

// init函数在包导入的时候自动执行
// 执行顺序: 全局声明 -> init() -> main()
// init函数没有参数, 没有返回值
func init() {
	fmt.Println("calc.init()")
}
