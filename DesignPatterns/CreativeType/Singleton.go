package CreativeType

import (
	"fmt"
	"sync"
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var instance *ChocolateBoiler
var once sync.Once

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

func (c *ChocolateBoiler) Fill() {
	if c.empty {
		c.empty = false
		fmt.Println("container is full")
	}
}

func (c *ChocolateBoiler) Drain() {
	if c.empty == false && c.boiled {
		c.empty = true
		c.boiled = false
		fmt.Println("poured into mold")
	}
}

func (c *ChocolateBoiler) Boil() {
	if c.empty == false && c.boiled == false {
		fmt.Println("chocolate boiled")
		c.boiled = true
	}
}
