package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//定义账号和密码
	acc string
	pwd string
	//定义一个变量接收用户的选择
	key string
	//声明一个变量，控制是否退出for循环
	loop bool //默认为真，可以一直循环
	//定义账户余额banance
	balance float64
	//定义每次收支的金额
	money float64
	//定义每次收支说明
	note string
	//定义一个变量，记录是否有收支行为
	flag bool //假定初始情况下是没有收支行为的
	//收支的详情使用字符串来记录，当有收支时，只需要对details进行拼接处理即可
	details string
}

//考虑到这个结构体首字母可能是小写的，所以编写一个“工厂模式”的方法，返回一个*FamilyAccount
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{ //要附初始值
		acc:     "",
		pwd:     "",
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说  明\t",
	}
}

//登录账号和密码
func (this *FamilyAccount) mima() {
	var n int = 3 //一共有登录次数3次
	var i int = 1 //登录几次
	choice1 := ""
	for {
		fmt.Println("请输入账号")
		fmt.Scanln(&this.acc)
		if len(this.acc) > 6 && len(this.acc) < 12 {
			fmt.Println("请输入密码")
			fmt.Scanln(&this.pwd)
			if len(this.pwd) > 0 && len(this.pwd) < 8 {
				fmt.Println("登陆成功")
				break
			}
		} else {
			fmt.Println("输入账号不正确，请重新输入")
			fmt.Println("剩余登录次数", n-i)
			if n-i == 0 {
				fmt.Println("输入密码不正确，是否退出y/n")
				fmt.Scanln(&choice1)
				if choice1 == "y" || choice1 == "n" {
					break
				} else {
					fmt.Println("请重新输入y/n")
				}
			}
			i++
		}
	}
	if choice1 == "y" {

		this.loop = false
	}

}

//1、将显示明细写成一个方法（封装）
func (this *FamilyAccount) showDetails() {
	fmt.Println("----------当前收支明细记录----------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("还没有收支记录，来添加一比吧~")
	}
}

//2、将登记收入也写成一个方法
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money //记录账户余额也会发生变化
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	//要将这个“收入”情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//3、支出
func (this *FamilyAccount) outcome() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	//这里需要一个必要的判断，如果这个支出的money比余额大
	if this.money > this.balance {
		fmt.Printf("余额不足，请重新输入，余额剩余为%v\n", this.balance)
	}
	this.balance -= this.money //记录账户余额也会发生变化
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	//要将这个“支出”情况，拼接到details变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	fmt.Println("----------登记完成----------")
	this.flag = true
}

//4、退出
func (this *FamilyAccount) quit() {
	fmt.Println("您确定要退出么?：y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("您输入的指令不对，请重新输入y/n")
		}
	}
	if choice == "y" {
		this.loop = false
	}
}
func (this *FamilyAccount) MainMenu() {
	//显示主菜单
	for {
		fmt.Println("\n----------智能记账本----------")
		fmt.Println("请先登录")
		this.mima()
		fmt.Println("          1 收支明细")
		fmt.Println("          2 登记收入")
		fmt.Println("          3 登记支出")
		fmt.Println("          4 退出")
		fmt.Print("          请选择（1-4）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.outcome()
		case "4":
			this.quit()

		default:
			fmt.Println("请输入正确的选项1-4")
			break
		}
		if this.loop == false {
			break
		}
	}
	fmt.Println("你已经退出该软件")

}
