// Package DesignPatterns contains design-principle and compound-pattern examples.
package DesignPatterns

import "fmt"

// ICar is the stable abstraction used to demonstrate the open-closed principle:
// clients depend on the car contract instead of concrete pricing details.
type ICar interface {
	GetName() string

	GetPrice() int
}

// BenzCar is the base implementation.
type BenzCar struct {
	name  string
	price int
}

func (b BenzCar) GetName() string {
	return b.name
}

func (b BenzCar) GetPrice() int {
	return b.price
}

// FinanceBenzCar extends BenzCar pricing without changing BenzCar itself.
type FinanceBenzCar struct {
	BenzCar
}

// GetPrice layers finance fees on top of the base car price.
func (b FinanceBenzCar) GetPrice() int {

	selfPrice := b.price
	var finance int
	if selfPrice >= 100 {
		finance = selfPrice + selfPrice*5/100
	} else if selfPrice >= 50 {
		finance = selfPrice + selfPrice*2/100
	} else {
		finance = selfPrice
	}
	return finance
}

// Girl is deliberately empty because this example focuses on responsibility
// boundaries rather than the data stored on each person.
type Girl struct {
}

// GroupLeader owns the collection and exposes the counting operation.
type GroupLeader struct {
	girls []Girl
}

func (g GroupLeader) CountGirls() {
	fmt.Println("The sum of girls is ", len(g.girls))
}

// Teacher depends on GroupLeader instead of reaching into the collection
// directly, illustrating the Law of Demeter.
type Teacher struct {
}

// Command asks the leader to count, keeping Teacher away from the internal list.
func (t Teacher) Command(leader GroupLeader) {
	leader.CountGirls()
}
