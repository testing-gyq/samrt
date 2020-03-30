package service

import (
	"go_code/chapter12项目/Customer/model"
)

//改CustomerService，完成对Customer的操作（增删改查）
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户，该字段后面还可以作为新用户的id+1
	CustomerNum int
}

// 编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，我们初始化一个客户
	customerService := &CustomerService{}
	customerService.CustomerNum = 1
	customers := model.NewCustomer(1, "张三", "男", 20, "112", "aaa@163.com")
	customerService.customers = append(customerService.customers, customers)
	return customerService

}

//返回客户信息切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户到customers切片中
func (this *CustomerService) Add(customer model.Customer) bool {
	//我们确定一个分配id的规则，就是添加的顺序
	this.CustomerNum++
	customer.ID = this.CustomerNum
	this.customers = append(this.customers, customer)
	return true
}

//根据ID删除客户，（从切片中删除）
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}

//编写一个方法更新用户信息
func (this *CustomerService) Update(customer model.Customer, id int) bool {
	index1 := this.FindById(id)
	if index1 == -1 {
		return false
	}
	this.customers[id-1] = customer
	return true
}

//根据id查找客户在切片中对应的下标,没有该客户返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	//遍历this.customers切片
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].ID == id {
			index = i
		}
	}
	return index
}
