package BehavioralType

import (
	"fmt"
	"testing"
)

func TestConcreteAggregate_Add(t *testing.T) {

	var (
		aggregate Aggregate
		iter      Iterator
	)
	aggregate = &ConcreteAggregate{docker: []interface{}{}}
	aggregate.Add("java")
	aggregate.Add("Golang")
	aggregate.Add("Python")

	iter = aggregate.CreateIterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
