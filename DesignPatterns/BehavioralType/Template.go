package BehavioralType

import "fmt"

type AbstractWork interface {
	GotoWork(work AbstractWork)
	Getup()
	Commute()
	Arrive()
}

type AbsClass struct{}

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
