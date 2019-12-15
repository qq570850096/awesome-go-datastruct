package CreativeType

import "testing"

func TestDirector(t *testing.T) {
	builder := &ConcreteBuilder{p: &Product{}}

	director := &Director{builder: builder}

	director.Construst()
}
