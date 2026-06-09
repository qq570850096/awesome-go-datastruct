package BehavioralType

// Iterator exposes collection traversal without exposing the collection's
// backing storage.
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// ConcreteIterator tracks traversal position over an Aggregate.
type ConcreteIterator struct {
	index int
	size  int
	con   Aggregate
}

func (c *ConcreteIterator) Next() interface{} {
	if c.HasNext() {
		res := c.con.GetElement(c.index)
		c.index++
		return res
	}
	return nil
}

func (c *ConcreteIterator) HasNext() bool {
	return c.index < c.size
}

type Aggregate interface {
	Add(obj interface{})
	CreateIterator() Iterator
	GetElement(index int) interface{}
	Size() int
}

// ConcreteAggregate stores elements and creates iterators over them.
type ConcreteAggregate struct {
	docker []interface{}
}

func (c *ConcreteAggregate) Add(obj interface{}) {
	c.docker = append(c.docker, obj)
}

func (c *ConcreteAggregate) CreateIterator() Iterator {
	// Capture the size at iterator creation time so this iterator has a stable
	// end position even if callers later add more items.
	return &ConcreteIterator{
		index: 0,
		size:  c.Size(),
		con:   c,
	}
}

func (c *ConcreteAggregate) GetElement(index int) interface{} {
	return c.docker[index]
}

func (c *ConcreteAggregate) Size() int {
	return len(c.docker)
}
