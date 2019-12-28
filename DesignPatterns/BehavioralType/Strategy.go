package BehavioralType

import "fmt"

type FlyBehavior interface {
	Fly()
}

type QuackBehavior interface {
	Quack()
}

type Duck struct {
	fly FlyBehavior
	quack QuackBehavior
}

func (d *Duck)Swim() {
	fmt.Println("鸭子游泳")
}

func (d *Duck) Display (behavior FlyBehavior,quackBehavior QuackBehavior) {
	behavior.Fly()
	quackBehavior.Quack()
}

type FlyWithWings struct {}

func (f *FlyWithWings) Fly ()  {
	fmt.Println("鸭子用翅膀飞")
}

type FlyNoWay struct {}

func (f *FlyNoWay) Fly ()  {
	fmt.Println("鸭子飞不起来")
}

type Quack struct {}

func (f *Quack) Quack ()  {
	fmt.Println("鸭子嘎嘎叫")
}

type Squeak struct {}

func (f *Squeak) Quack ()  {
	fmt.Println("鸭子咔咔叫")
}

type Mute struct {}

func (f *Mute) Quack ()  {
	fmt.Println("鸭子不能叫")
}

type ReadHead struct {
	*Duck
	fly *FlyWithWings
	quack *Quack
}

func (r *ReadHead) Display ()  {
	r.Swim()
	r.Duck.Display(r.fly, r.quack)
}

type Wooden struct {
	*Duck
	fly *FlyNoWay
	quack *Mute
}

func (r *Wooden) Display ()  {
	r.Swim()
	r.Duck.Display(r.fly,r.quack)
}

type Mallard struct {
	*Duck
	fly *FlyWithWings
	quack *Quack
}

func (m *Mallard) Display ()  {
	m.Swim()
	m.Duck.Display(m.fly, m.quack)
}

type Rubber struct {
	*Duck
	fly *FlyNoWay
	quack *Squeak
}

func (r *Rubber) Display ()  {
	r.Swim()
	r.Duck.Display(r.fly, r.quack)
}