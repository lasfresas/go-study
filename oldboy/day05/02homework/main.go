package main

import (
	"fmt"
	"os"
)

/*
	结构体版学生管理系统
	写一个系统能够查看、新增、修改学生、删除学生
	1. 它保存了一些数据 ---> 结构体的字段
	2. 它有4个功能 ---> 结构体的方法
*/

var smr studentMgr // 声明一个全局的变量学生管理对象：smr

// 菜单
func showMenu() {
	fmt.Println("-----welcome sms!-----")
	fmt.Println(`
	1. 查看学生
	2. 增加学生
	3. 修改学生
	4. 删除学生
	5. 退出
	`)
}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		// 等待用户输入选项
		fmt.Print("请输入你的选项:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的是:%d\n", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudents()
		case 3:
			smr.editStudents()
		case 4:
			smr.deleteStudents()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("滚~")
		}
	}
}
