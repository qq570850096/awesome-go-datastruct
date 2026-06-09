package BehavioralType

import "fmt"

// AbstractWork defines the template-method steps. Getup and Arrive are fixed
// in AbsClass, while Commute is supplied by concrete work routes.
type AbstractWork interface {
	GotoWork(work AbstractWork)
	Getup()
	Commute()
	Arrive()
}

// AbsClass owns the invariant algorithm skeleton.
type AbsClass struct{}

// GotoWork is the template method: fixed steps wrap the variable commute step.
func (a AbsClass) GotoWork(work AbstractWork) {
	a.Getup()
	work.Commute()
	a.Arrive()
}

func (a AbsClass) Getup() {
	fmt.Println("1. wake up")
}

func (a AbsClass) Commute() {}

func (a AbsClass) Arrive() {
	fmt.Println("3. arrive")
}

type DriveToWork struct {
	AbsClass
}

func (d *DriveToWork) Commute() {
	fmt.Println("2. drive to work")
}

func (d *DriveToWork) GotoWork(work AbstractWork) {
	d.AbsClass.GotoWork(d)
}

type BusToWork struct {
	AbsClass
}

func (d *BusToWork) Commute() {
	fmt.Println("2. take bus to work")
}

func (d *BusToWork) GotoWork(work AbstractWork) {
	d.AbsClass.GotoWork(d)
}
