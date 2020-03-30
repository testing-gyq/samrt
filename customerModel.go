package model

import (
	"fmt"
)

//声明一个结构体表示客户信息
type Customer struct {
	ID     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个Cuctomer实例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//第二种创建Cuctomer实例的方法，不带ID
func NewCustomer2(name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户信息,格式化输出
func (this Customer) GetInfo() string {
	Info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		this.ID, this.Name, this.Gender, this.Age, this.Phone, this.Email)
	return Info
}
