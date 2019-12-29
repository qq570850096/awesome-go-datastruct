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

// 第一次需求变动，现在农场多了一些鹅，鹅也想quack()但是鹅只有creak()
// 采用适配器模式
// 适配模式
type Creak interface {
	Creak()
	QuackObservable
}
// 被适配的类
type Goose struct {
	observable QuackObservable
}

func (g *Goose) Creak() {
	fmt.Println("鹅喳喳叫")
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
// 适配器提供了统一的方法，这样鹅也可以quack了
func (g *GooseAdapter) quack() {
	g.Creak.Creak()
}

// 第二次变动，统计叫声
// 采用修饰器模式
type QuackCounter struct {
	q QuackAble
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
	fmt.Println("鸭子第",q.numberOfQuacks,"次叫")
}

// 第三次需求变更为了防止调用者忘了建立叫声统计类
// 我们这里使用抽象工厂,开闭原则中对扩展是开放的，所以我们将抽象工厂也扩展到修饰器中
type AbsDuckFactory interface {
	CreateMallardDuck() QuackAble
	CreateRedheadDuck() QuackAble
	CreateDuckCall() QuackAble
	CreateRubber() QuackAble
}

// 鸭子工厂
type DuckFactory struct {

}

func (q *DuckFactory) CreateMallardDuck() QuackAble {
	return &QuackCounter{&MallardDuck{observable:&ObservableAssist{
		list:            []Observer{},
		quackObservable: &MallardDuck{},
	}},0}
}

func (q *DuckFactory) CreateRedheadDuck() QuackAble {
	return &QuackCounter{
		q:              &RedheadDuck{observable:&ObservableAssist{
			list:            []Observer{},
			quackObservable: &RedheadDuck{},
		}},
		numberOfQuacks: 0,
	}
}

func (q *DuckFactory) CreateDuckCall() QuackAble {
	return &QuackCounter{
		q:              &DuckCall{observable:&ObservableAssist{
			list:            []Observer{},
			quackObservable: &DuckCall{},
		}},
		numberOfQuacks: 0,
	}
}

func (q *DuckFactory) CreateRubber() QuackAble {
	return &QuackCounter{
		q:              &RedheadDuck{observable:&ObservableAssist{
			list:            []Observer{},
			quackObservable: &RubberDuck{},
		}},
		numberOfQuacks: 0,
	}
}

// 第四次需求变更，我们需要管理那些鸭子和鹅
// 这时候使用组合模式
// 组合模式接口
type FlockDuck interface {
	Add (q QuackAble)
	Remove(q QuackAble)
}

type Flock struct {
	qs []QuackAble
}

func (f *Flock) RegisterObserver(observer Observer) {
	for _,v := range f.qs {
		v.RegisterObserver(observer)
	}
}

func (f *Flock) RemoveObserver(observer Observer) {
	for _,v := range f.qs {
		v.RemoveObserver(observer)
	}
}

func (f *Flock) NotifyObservers() {
	for _,v := range f.qs {
		v.NotifyObservers()
	}
}

func (f *Flock) quack() {
	for _,v := range f.qs {
		v.quack()
	}
}

func (f *Flock) Add (q QuackAble)  {
	f.qs = append(f.qs,q)
}

func (f *Flock) Remove (q QuackAble)  {
	for i,v := range f.qs {
		if v == q {
			// 删除切片中的第i个元素
			f.qs = append(f.qs[:i],f.qs[i+1:]...)
		}
	}
}

// 第五次需求变更，我们想要做一个提醒，当任何一个鸭子或鹅叫的时候
// 毫无疑问，这时候应该采取观察者模式
// 订阅接收者接口
type Observer interface {
	Update(observable QuackObservable)
}
// 订阅发布者接口
type QuackObservable interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}
// 具体的订阅接收者
type DuckDoctor struct {

}

func (d DuckDoctor) Update(observable QuackObservable) {
	fmt.Printf("DuckDoctor 观察到的鸭子对象为%T\n",observable)
}

// 具体的发布者
type ObservableAssist struct {
	list []Observer
	quackObservable QuackObservable
}

func (o *ObservableAssist) RegisterObserver(observer Observer) {
	o.list = append(o.list,observer)
}

func (o *ObservableAssist) RemoveObserver(observer Observer) {
	for i,v := range o.list {
		if v == observer {
			o.list = append(o.list[:i],o.list[i+1:]...)
		}
	}
}

func (o *ObservableAssist) NotifyObservers() {
	for _,v := range o.list {
		fmt.Printf("订阅者%T收到如下信息\n",v)
		v.Update(o.quackObservable)
	}
}