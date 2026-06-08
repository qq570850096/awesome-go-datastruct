package DesignPatterns

import "fmt"

type QuackAble interface {
	quack()
	QuackObservable
}

type MallardDuck struct {
	observable QuackObservable
}

func (m *MallardDuck) RegisterObserver(observer Observer) {
	m.observable.RegisterObserver(observer)
}

func (m *MallardDuck) RemoveObserver(observer Observer) {
	m.observable.RemoveObserver(observer)
}

func (m *MallardDuck) NotifyObservers() {
	m.observable.NotifyObservers()
}

func (m *MallardDuck) quack() {
	fmt.Println("mallardDuck")
}

type RedheadDuck struct {
	observable QuackObservable
}

func (r *RedheadDuck) RegisterObserver(observer Observer) {
	r.observable.RegisterObserver(observer)
}

func (r *RedheadDuck) RemoveObserver(observer Observer) {
	r.observable.RemoveObserver(observer)
}

func (r *RedheadDuck) NotifyObservers() {
	r.observable.NotifyObservers()
}

func (r *RedheadDuck) quack() {
	fmt.Println("RedheadDuck")
}

type DuckCall struct {
	observable QuackObservable
}

func (d *DuckCall) RegisterObserver(observer Observer) {
	d.observable.RegisterObserver(observer)
}

func (d *DuckCall) RemoveObserver(observer Observer) {
	d.observable.RemoveObserver(observer)
}

func (d *DuckCall) NotifyObservers() {
	d.observable.NotifyObservers()
}

func (d *DuckCall) quack() {
	fmt.Println("DuckCall")
}

type RubberDuck struct {
	observable QuackObservable
}

func (r *RubberDuck) RegisterObserver(observer Observer) {
	r.observable.RegisterObserver(observer)
}

func (r *RubberDuck) RemoveObserver(observer Observer) {
	r.observable.RemoveObserver(observer)
}

func (r *RubberDuck) NotifyObservers() {
	r.observable.NotifyObservers()
}

func (r *RubberDuck) quack() {
	fmt.Println("RubberDuck")
}

type Creak interface {
	Creak()
	QuackObservable
}

type Goose struct {
	observable QuackObservable
}

func (g *Goose) Creak() {
	fmt.Println("goose honks")
}

func (g *Goose) RegisterObserver(observer Observer) {
	g.observable.RegisterObserver(observer)
}

func (g *Goose) RemoveObserver(observer Observer) {
	g.observable.RemoveObserver(observer)
}

func (g *Goose) NotifyObservers() {
	g.observable.NotifyObservers()
}

type GooseAdapter struct {
	Creak
}

func (g *GooseAdapter) quack() {
	g.Creak.Creak()
}

type QuackCounter struct {
	q              QuackAble
	numberOfQuacks int
}

func (q *QuackCounter) RegisterObserver(observer Observer) {
	q.q.RegisterObserver(observer)
}

func (q *QuackCounter) RemoveObserver(observer Observer) {
	q.q.RemoveObserver(observer)
}

func (q *QuackCounter) NotifyObservers() {
	q.q.NotifyObservers()
}

func (q *QuackCounter) quack() {
	q.q.quack()
	q.numberOfQuacks++
	fmt.Println("duck quack #", q.numberOfQuacks, "time")
}

type AbsDuckFactory interface {
	CreateMallardDuck() QuackAble
	CreateRedheadDuck() QuackAble
	CreateDuckCall() QuackAble
	CreateRubber() QuackAble
}

type DuckFactory struct {
}

func (q *DuckFactory) CreateMallardDuck() QuackAble {
	return &QuackCounter{&MallardDuck{observable: &ObservableAssist{
		list:            []Observer{},
		quackObservable: &MallardDuck{},
	}}, 0}
}

func (q *DuckFactory) CreateRedheadDuck() QuackAble {
	return &QuackCounter{
		q: &RedheadDuck{observable: &ObservableAssist{
			list:            []Observer{},
			quackObservable: &RedheadDuck{},
		}},
		numberOfQuacks: 0,
	}
}

func (q *DuckFactory) CreateDuckCall() QuackAble {
	return &QuackCounter{
		q: &DuckCall{observable: &ObservableAssist{
			list:            []Observer{},
			quackObservable: &DuckCall{},
		}},
		numberOfQuacks: 0,
	}
}

func (q *DuckFactory) CreateRubber() QuackAble {
	return &QuackCounter{
		q: &RedheadDuck{observable: &ObservableAssist{
			list:            []Observer{},
			quackObservable: &RubberDuck{},
		}},
		numberOfQuacks: 0,
	}
}

type FlockDuck interface {
	Add(q QuackAble)
	Remove(q QuackAble)
}

type Flock struct {
	qs []QuackAble
}

func (f *Flock) RegisterObserver(observer Observer) {
	for _, v := range f.qs {
		v.RegisterObserver(observer)
	}
}

func (f *Flock) RemoveObserver(observer Observer) {
	for _, v := range f.qs {
		v.RemoveObserver(observer)
	}
}

func (f *Flock) NotifyObservers() {
	for _, v := range f.qs {
		v.NotifyObservers()
	}
}

func (f *Flock) quack() {
	for _, v := range f.qs {
		v.quack()
	}
}

func (f *Flock) Add(q QuackAble) {
	f.qs = append(f.qs, q)
}

func (f *Flock) Remove(q QuackAble) {
	for i, v := range f.qs {
		if v == q {

			f.qs = append(f.qs[:i], f.qs[i+1:]...)
		}
	}
}

type Observer interface {
	Update(observable QuackObservable)
}

type QuackObservable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type DuckDoctor struct {
}

func (d DuckDoctor) Update(observable QuackObservable) {
	fmt.Printf("DuckDoctor observed duck object %T\n", observable)
}

type ObservableAssist struct {
	list            []Observer
	quackObservable QuackObservable
}

func (o *ObservableAssist) RegisterObserver(observer Observer) {
	o.list = append(o.list, observer)
}

func (o *ObservableAssist) RemoveObserver(observer Observer) {
	for i, v := range o.list {
		if v == observer {
			o.list = append(o.list[:i], o.list[i+1:]...)
		}
	}
}

func (o *ObservableAssist) NotifyObservers() {
	for _, v := range o.list {
		fmt.Printf("subscriber %T received the following information\n", v)
		v.Update(o.quackObservable)
	}
}
