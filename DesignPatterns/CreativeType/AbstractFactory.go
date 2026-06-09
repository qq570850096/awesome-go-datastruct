// Package CreativeType contains examples of creational design patterns.
package CreativeType

import "fmt"

// Factory is the abstract factory role. It creates a compatible family of
// products without exposing the concrete product types to the client.
type Factory interface {
	NewTV() Television
	NewRefrigerator() Refrigerator
}

// Television is one product interface in the appliance family.
type Television interface {
	DoSomething()
}

// Refrigerator is another product interface in the same appliance family.
type Refrigerator interface {
	DoSomething()
}

// TCLTV is a concrete TV product created by TCLFactory.
type TCLTV struct {
}

func (TCLTV) DoSomething() {
	fmt.Println("TCL TV is doing something")
}

// TCLRef is a concrete refrigerator product created by TCLFactory.
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

// MediaRef is a concrete refrigerator product created by MediaFactory.
type MediaRef struct{}

func (MediaRef) DoSomething() {
	fmt.Println("Midea air conditioner is doing something")
}

type TCLFactory struct {
}

// NewTV returns the TV variant that belongs to the TCL product family.
func (TCLFactory) NewTV() Television {
	return TCLTV{}
}

// NewRefrigerator returns the refrigerator variant that belongs to the TCL family.
func (TCLFactory) NewRefrigerator() Refrigerator {
	return TCLRef{}
}

// MediaFactory creates the Midea product family.
type MediaFactory struct {
}

// NewTV returns the TV variant that belongs to the Midea product family.
func (MediaFactory) NewTV() Television {
	return MediaTV{}
}

// NewRefrigerator returns the refrigerator variant that belongs to the Midea family.
func (MediaFactory) NewRefrigerator() Refrigerator {
	return MediaRef{}
}
