package DesignPatterns

import (
	"testing"
)

func TestDuckCall_RemoveObserver(t *testing.T) {
	var (
		factory   AbsDuckFactory
		quackable QuackAble
		creak     Creak
		flok      *Flock
		observer  Observer
	)
	t.Log("test abstract factory")
	factory = &DuckFactory{}
	quackable = factory.CreateMallardDuck()
	quackable.quack()
	quackable = factory.CreateRedheadDuck()
	quackable.quack()
	quackable = factory.CreateDuckCall()
	quackable.quack()
	quackable = factory.CreateRubber()
	quackable.quack()
	t.Log("test adapter")
	creak = &Goose{}
	adapter := &GooseAdapter{creak}
	adapter.quack()
	t.Log("test composite")
	flok = &Flock{qs: []QuackAble{}}
	flok.Add(quackable)
	flok.quack()
	t.Log("test observer")
	observer = &DuckDoctor{}

	quackable.RegisterObserver(observer)

	flok.NotifyObservers()
	quackable = factory.CreateRedheadDuck()
	quackable.RegisterObserver(observer)

	flok.Add(quackable)
	flok.NotifyObservers()
}
