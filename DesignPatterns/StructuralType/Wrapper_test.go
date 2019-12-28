package StructuralType

import "testing"

func TestAdapter_TargetMethod1(t *testing.T) {
	adapter := Adapter{}

	adapter.Quack()
	adapter.Gobble()
	adapter.Fly()
	adapter.TurkeyFly()
}
