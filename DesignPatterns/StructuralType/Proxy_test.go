package StructuralType

import "testing"

func TestBuyer_BuyTicket(t *testing.T) {
	buyer := &Buyer{name: "123"}
	proxy := BuyerProxy{b: buyer}
	proxy.Login("678","345")
	proxy.BuyTicket()
}
