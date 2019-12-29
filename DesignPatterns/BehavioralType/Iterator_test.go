package BehavioralType

import (
	"fmt"
	"testing"
)

func TestConcreteAggregate_Add(t *testing.T) {
	// 定义聚族对象
	var (
		aggregate Aggregate
		iter Iterator
	)
	aggregate = &ConcreteAggregate{docker: []interface{}{}}
	aggregate.Add("java")
	aggregate.Add("Golang")
	aggregate.Add("Python")
	// 遍历
	iter = aggregate.CreateIterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
