package BehavioralType

import "testing"

func TestReadHead_Display(t *testing.T) {
	flynoway := &FlyNoWay{}
	flayWihtwings := &FlyWithWings{}
	quack := &Quack{}
	sqeak := &Squeak{}
	mute := &Mute{}
	duck := ReadHead{
		Duck:  &Duck{},
		fly:   flayWihtwings,
		quack: quack,
	}
	duck.Display()
	mallard := Mallard {
		Duck:  &Duck{},
		fly:   flayWihtwings,
		quack: quack,
	}
	mallard.Display()
	rub := Rubber {
		Duck:  &Duck{},
		fly:   flynoway,
		quack: sqeak,
	}
	rub.Display()
	wooden := Wooden{
		Duck:  &Duck{},
		fly:   flynoway,
		quack: mute,
	}
	wooden.Display()
}