package BehavioralType

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

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

type ConcreteAggregate struct {
	docker []interface{}
}

func (c *ConcreteAggregate) Add(obj interface{}) {
	c.docker = append(c.docker, obj)
}

func (c *ConcreteAggregate) CreateIterator() Iterator {
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
