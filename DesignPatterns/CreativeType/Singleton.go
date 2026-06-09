package CreativeType

import (
	"fmt"
	"sync"
)

// ChocolateBoiler is the singleton-managed object. Its state transitions mimic
// the Head First Design Patterns chocolate boiler example.
type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

// instance is package-level shared state guarded by once.
var instance *ChocolateBoiler
var once sync.Once

// GetInstance returns the only ChocolateBoiler instance. sync.Once guarantees
// that initialization runs once even when multiple goroutines call it.
func GetInstance() *ChocolateBoiler {

	once.Do(func() {
		fmt.Println("pot created")
		instance = &ChocolateBoiler{true, false}
	})
	return instance
}

func (c *ChocolateBoiler) IsEmpty() bool {
	return c.empty
}

func (c *ChocolateBoiler) IsBoiled() bool {
	return c.boiled
}

// Fill is valid only when the boiler is empty.
func (c *ChocolateBoiler) Fill() {
	if c.empty {
		c.empty = false
		fmt.Println("container is full")
	}
}

// Drain is valid only after the boiler has been filled and boiled.
func (c *ChocolateBoiler) Drain() {
	if c.empty == false && c.boiled {
		c.empty = true
		c.boiled = false
		fmt.Println("poured into mold")
	}
}

// Boil moves a filled boiler into the boiled state.
func (c *ChocolateBoiler) Boil() {
	if c.empty == false && c.boiled == false {
		fmt.Println("chocolate boiled")
		c.boiled = true
	}
}
