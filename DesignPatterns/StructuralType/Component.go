package StructuralType

import "fmt"

// MenuComponent is the component interface shared by leaves and composites.
type MenuComponent interface {
	Add(menuComponent MenuComponent)
	Remove(menuComponent MenuComponent)

	GetName() string
	GetPrice() float64
	GetDescription() string
	IsVegetarian() bool
	CreateIterator()
	Display(depth int)
}

// Leaf is a terminal menu item. Child-management operations panic because a
// leaf cannot contain other menu components.
type Leaf struct {
	name        string
	vegetarian  bool
	description string
	price       float64
}

func (l *Leaf) Add(menuComponent MenuComponent) {
	panic("Leaf Node can not add")
}

func (l *Leaf) Remove(menuComponent MenuComponent) {
	panic("Leaf Node can not remove")
}

func (l *Leaf) GetName() string {
	return l.name
}

func (l *Leaf) GetPrice() float64 {
	return l.price
}

func (l *Leaf) IsVegetarian() bool {
	return l.vegetarian
}

func (l *Leaf) GetDescription() string {
	return l.description
}

func (l *Leaf) CreateIterator() {
	panic("implement me")
}

func (l *Leaf) Display(depth int) {

	for i := 0; i < depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(l.Name())
}

func (l *Leaf) Name() string {
	return l.name
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

// Composite groups menu components and forwards tree operations to children.
type Composite struct {
	name        string
	description string
	arr         []MenuComponent
}

func (c *Composite) GetName() string {
	return c.name
}

func (c *Composite) GetPrice() float64 {
	panic("It is not an item.")
}

func (c *Composite) GetDescription() string {
	return c.description
}

func (c *Composite) IsVegetarian() bool {
	panic("implement me")
}

func (c *Composite) CreateIterator() {
	panic("implement me")
}

func (c *Composite) Add(MenuComponent MenuComponent) {
	c.arr = append(c.arr, MenuComponent)
}

func (c *Composite) Remove(MenuComponent MenuComponent) {
	for i, v := range c.arr {
		if v == MenuComponent {

			c.arr = append(c.arr[:i], c.arr[i+1:]...)
		}
	}
}

func (c *Composite) Display(depth int) {

	for i := 0; i < depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(c.GetName())

	for _, com := range c.arr {
		com.Display(depth + 1)
	}
}
