package CreativeType

import "fmt"

type Factory interface {
	NewTV() Television
	NewRefrigerator() Refrigerator
}

type Television interface {
	DoSomething()
}

type Refrigerator interface {
	DoSomething()
}

type TCLTV struct {
}

func (TCLTV) DoSomething() {
	fmt.Println("TCL TV is doing something")
}

type TCLRef struct {
}

func (TCLRef) DoSomething() {
	fmt.Println("TCL air conditioner is doing something")
}

type MediaTV struct {
}

func (MediaTV) DoSomething() {
	fmt.Println("Midea TV is doing something")
}

type MediaRef struct{}

func (MediaRef) DoSomething() {
	fmt.Println("Midea air conditioner is doing something")
}

type TCLFactory struct {
}

func (TCLFactory) NewTV() Television {
	return TCLTV{}
}

func (TCLFactory) NewRefrigerator() Refrigerator {
	return TCLRef{}
}

type MediaFactory struct {
}

func (MediaFactory) NewTV() Television {
	return MediaTV{}
}

func (MediaFactory) NewRefrigerator() Refrigerator {
	return MediaRef{}
}
