package main

import (
	"fmt"
	"go_code/chapter12项目/oop实现记账/utils"
)

func main() {
	fmt.Println("这是用面向对象的方法实现的")
	//创建一个NewFamilyAccoun指针结构体变量，引用MainMenu方法
	utils.NewFamilyAccount().MainMenu()

}
