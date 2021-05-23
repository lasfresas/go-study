package main

import "fmt"

// 标识符：变量名，类型名，函数名，方法名
// Go中如果标识符首字母是大写的，就表示对外可见（暴露的，公有的）

type dog struct {
	name string
}

// 构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// Go语言中的方法（Method）是一种作用于特定类型变量的函数。
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~\n", d.name)
}

func main() {
	d1 := newDog("皮皮")
	d1.wang()
}
