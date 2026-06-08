package BehavioralType

import (
	"fmt"
	"testing"
)

func TestCaretaker_GetMemento(t *testing.T) {

	o := &Originator{state: "hello"}
	fmt.Println("current state:", o.GetState())

	c := new(Caretaker)
	c.SetMemento(o.CreateMemento())

	o.SetState("world")
	fmt.Println("changed state:", o.GetState())

	o.SetState(c.GetMemento().GetState())
	fmt.Println("restored state", o.GetState())
}
