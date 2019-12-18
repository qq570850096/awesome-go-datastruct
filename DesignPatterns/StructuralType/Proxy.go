package StructuralType

import "fmt"

type IBuyer interface {
	Login(username,password string)
	BuyTicket()
}

type BuyerProxy struct {
	b *Buyer
}

func (b *BuyerProxy)Login(username,password string)  {
	b.b.Login(username,password)
}
func (b *BuyerProxy)BuyTicket()  {
	before()
	b.b.BuyTicket()
	after()
}

func before() {
	fmt.Println("准备定时任务，开始刷票")
}

func after() {
	fmt.Println("刷票成功，短信通知用户")
}

type Buyer struct {
	name string
}

func (b *Buyer)Login(username,password string)  {
	fmt.Println(b.name,"使用",username,"登陆成功")
}

func (b *Buyer)BuyTicket()  {
	fmt.Println(b.name,"购票成功")
}


