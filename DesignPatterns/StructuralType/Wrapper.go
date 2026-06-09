package StructuralType

import "fmt"

// Duck is the target interface expected by the client.
type Duck interface {
	Quack()
	Fly()
}

// Turkey is the incompatible interface that needs adapting.
type Turkey interface {
	Gobble()
	TurkeyFly()
}

// Adaptee is the existing turkey implementation.
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

// Quack translates the target duck call into behavior compatible with the
// wrapped turkey example.
func (a Adapter) Quack() {
	fmt.Println("duck quacks")
}

// Fly satisfies the Duck target interface.
func (a Adapter) Fly() {
	fmt.Println("duck quacks")
}
