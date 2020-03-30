package main

import (
	"fmt"
	"go_code/chapter12项目/Customer/model"
	"go_code/chapter12项目/Customer/service"
)

type customerView struct {
	key             string                   //接收用户输入
	loop            bool                     //表示是否循环主菜单
	customerService *service.CustomerService //增加一个字段customerService
}

//显示所有用户信息
func (this *customerView) list() {
	//首先，获取当前所有用户的信息
	customers := this.customerService.List()
	fmt.Println("\n----------客户列表----------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("----------客户列表完成----------\n")
}

//得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意：id号，没有让用户输入，因为id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("----------添加成功----------")
	} else {
		fmt.Println("----------添加失败----------")
	}
}

//得到用户的输入，并删除该ID对应的用户
func (this *customerView) delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("请选择待删除客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除的操作
	}
	fmt.Println("确认是否删除（Y/N）：")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "N" {
			if this.customerService.Delete(id) {
				fmt.Println("----------删除成功----------")
			} else {
				fmt.Println("删除失败，您输入的id号不存在")
			}
			break
		} else {
			fmt.Println("您的输入有误，请重新输入：（Y/N）")
		}
	}
	//如果ID不等于-1，那么就调用方法
}

//修改客户信息
func (this *customerView) update() {
	fmt.Println("----------修改客户----------")
	fmt.Println("请选择待修改客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id, name, gender, age, phone, email)
	if this.customerService.Update(customer, id) {

		fmt.Println("----------修改成功----------")
	} else {
		fmt.Println("----------修改失败----------")
	}
}

//退出确认（Y/N）
func (this *customerView) queRen() {
	fmt.Println("确认是否退出该系统：（Y/N）")

	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "N" {
			if this.key == "Y" {
				break
			}
		} else {
			fmt.Println("您输入的信息有误，请重新打呼入(Y/N)")
		}
	}
	if this.key == "Y" {
		this.loop = false
	}
}

//显示主菜单
func (this *customerView) mainView() {

	for this.loop {
		fmt.Println("----------客户信息管理软件----------")
		fmt.Println("          1 添加客户")
		fmt.Println("          2 修改客户")
		fmt.Println("          3 删除客户")
		fmt.Println("          4 客户列表")
		fmt.Println("          5 退   出")
		fmt.Print("请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.queRen()
		default:
			fmt.Println("您输入的信息有误，请重新输入(1-5)")
		}
	}
	fmt.Println("您已经退出了客户关系管理系统")
}

func main() {
	//这里完成对customerView结构体的customerService,
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//对customerView结构体的customerService字段初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainView()
}
