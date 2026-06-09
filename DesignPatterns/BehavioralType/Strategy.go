package BehavioralType

import "fmt"

// FlyBehavior is the strategy interface for flying.
type FlyBehavior interface {
	Fly()
}

// QuackBehavior is the strategy interface for making sounds.
type QuackBehavior interface {
	Quack()
}

// Duck is the context that delegates variable behavior to strategy objects.
type Duck struct {
	fly   FlyBehavior
	quack QuackBehavior
}

func (d *Duck) Swim() {
	fmt.Println("duck swims")
}

func (d *Duck) Display(behavior FlyBehavior, quackBehavior QuackBehavior) {
	behavior.Fly()
	quackBehavior.Quack()
}

// FlyWithWings is a concrete flying strategy.
type FlyWithWings struct{}

func (f *FlyWithWings) Fly() {
	fmt.Println("duck flies with wings")
}

type FlyNoWay struct{}

func (f *FlyNoWay) Fly() {
	fmt.Println("duck cannot fly")
}

type Quack struct{}

func (f *Quack) Quack() {
	fmt.Println("duck quacks")
}

type Squeak struct{}

func (f *Squeak) Quack() {
	fmt.Println("duck squeaks")
}

type Mute struct{}

func (f *Mute) Quack() {
	fmt.Println("duck cannot quack")
}

type ReadHead struct {
	*Duck
	fly   *FlyWithWings
	quack *Quack
}

// Display combines fixed duck behavior with the readhead duck strategies.
func (r *ReadHead) Display() {
	r.Swim()
	r.Duck.Display(r.fly, r.quack)
}

type Wooden struct {
	*Duck
	fly   *FlyNoWay
	quack *Mute
}

func (r *Wooden) Display() {
	r.Swim()
	r.Duck.Display(r.fly, r.quack)
}

type Mallard struct {
	*Duck
	fly   *FlyWithWings
	quack *Quack
}

func (m *Mallard) Display() {
	m.Swim()
	m.Duck.Display(m.fly, m.quack)
}

type Rubber struct {
	*Duck
	fly   *FlyNoWay
	quack *Squeak
}

func (r *Rubber) Display() {
	r.Swim()
	r.Duck.Display(r.fly, r.quack)
}
