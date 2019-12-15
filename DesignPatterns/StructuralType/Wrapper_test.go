package StructuralType

import "testing"

func TestAdapter_TargetMethod1(t *testing.T) {
	adapter := Adapter{}

	adapter.TargetMethod1()
	adapter.TargetMethod2()
}
