package main

import "fmt"

type common interface {
	login() 
	logout()
}

type user interface {
	common
}

type vip_interface interface {
	ConsulTation()
	common
}

type user_1 struct {
	name string
}

func (u user_1) login() {
	fmt.Println("User login")
}

func (u user_1) logout() {
	fmt.Println("User logout")
}

func (u user_1) ConsulTation() {
	fmt.Println("VIP User consultation")
}

func main() {
	var user user_1
	user = user_1{"黄昌昊"}
	var vipuser vip_interface
	vipuser = user
	vipuser.login()
}