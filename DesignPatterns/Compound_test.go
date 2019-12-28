package DesignPatterns

import (
	"testing"
)

// 复合模式测试代码
func TestDuckCall_RemoveObserver(t *testing.T) {
	var (
		factory AbsDuckFactory
		quackable QuackAble
		creak Creak
		flok *Flock
		observer Observer
	)
	t.Log("测试抽象工厂")
	factory = &DuckFactory{}
	quackable = factory.CreateMallardDuck()
	quackable.quack()
	quackable = factory.CreateRedheadDuck()
	quackable.quack()
	quackable = factory.CreateDuckCall()
	quackable.quack()
	quackable = factory.CreateRubber()
	quackable.quack()
	t.Log("测试适配器")
	creak = &Goose{}
	adapter := &GooseAdapter{creak}
	adapter.quack()
	t.Log("测试组合模式")
	flok = &Flock{qs: []QuackAble{}}
	flok.Add(quackable)
	flok.quack()
	t.Log("测试观察者模式")
	observer = &DuckDoctor{}

	quackable.RegisterObserver(observer)
	//quackable.NotifyObservers()
	flok.NotifyObservers()
	quackable = factory.CreateRedheadDuck()
	quackable.RegisterObserver(observer)
	//quackable.NotifyObservers()
	flok.Add(quackable)
	flok.NotifyObservers()
}
