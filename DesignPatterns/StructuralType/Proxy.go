package StructuralType

import "fmt"

// IBuyer is the subject interface shared by the real buyer and the proxy.
type IBuyer interface {
	Login(username, password string)
	BuyTicket()
}

// BuyerProxy controls access to Buyer by adding behavior before and after the
// real ticket purchase.
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

// Buyer is the real subject that performs login and ticket purchasing.
type Buyer struct {
	name string
}

func (b *Buyer) Login(username, password string) {
	fmt.Println(b.name, "uses", username, "login succeeded")
}

func (b *Buyer) BuyTicket() {
	fmt.Println(b.name, "ticket purchase succeeded")
}
