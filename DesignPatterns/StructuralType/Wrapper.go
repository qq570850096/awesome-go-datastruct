package StructuralType

import "fmt"

type Duck interface {
	Quack()
	Fly()
}

type Turkey interface {
	Gobble()
	TurkeyFly()
}

type Adaptee struct{}

func (a Adaptee) Gobble() {
	fmt.Println("turkey gobbles")
}

func (a Adaptee) TurkeyFly() {
	fmt.Println("turkey flies")
}

type Adapter struct {
	Adaptee
}

func (a Adapter) Quack() {
	fmt.Println("duck quacks")
}

func (a Adapter) Fly() {
	fmt.Println("duck quacks")
}
