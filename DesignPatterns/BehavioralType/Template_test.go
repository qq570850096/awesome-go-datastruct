package BehavioralType

import "testing"

func TestAbsClass_GotoWork(t *testing.T) {
	var (
		work AbstractWork
	)
	work = &BusToWork{AbsClass{}}
	work.GotoWork(work)
	work = &DriveToWork{AbsClass{}}
	work.GotoWork(work)

}
