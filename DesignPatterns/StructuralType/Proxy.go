package StructuralType

import "fmt"

type IBuyer interface {
	Login(username, password string)
	BuyTicket()
}

type BuyerProxy struct {
	b *Buyer
}

func (b *BuyerProxy) Login(username, password string) {
	b.b.Login(username, password)
}
func (b *BuyerProxy) BuyTicket() {
	before()
	b.b.BuyTicket()
	after()
}

func before() {
	fmt.Println("prepare scheduled task and start ticket polling")
}

func after() {
	fmt.Println("ticket acquired, notify user by SMS")
}

type Buyer struct {
	name string
}

func (b *Buyer) Login(username, password string) {
	fmt.Println(b.name, "uses", username, "login succeeded")
}

func (b *Buyer) BuyTicket() {
	fmt.Println(b.name, "ticket purchase succeeded")
}
