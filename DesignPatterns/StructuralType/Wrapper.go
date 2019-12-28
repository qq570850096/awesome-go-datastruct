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

type Adaptee struct {}

func (a Adaptee) Gobble() {
	fmt.Println("火鸡咕咕叫")
}

func (a Adaptee) TurkeyFly() {
	fmt.Println("火鸡起飞")
}

type Adapter struct {
	Adaptee
}

func (a Adapter) Quack() {
	fmt.Println("鸭子嘎嘎叫")
}

func (a Adapter) Fly() {
	fmt.Println("鸭子嘎嘎叫")
}


