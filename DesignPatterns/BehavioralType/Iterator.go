package BehavioralType


// 抽象迭代器
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// 具体迭代器
type ConcreteIterator struct {
	index int
	size int
	con Aggregate
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

// 抽象聚集
type Aggregate interface {
	Add(obj interface{})
	CreateIterator() Iterator
	GetElement(index int) interface{}
	Size() int
}

// 具体聚集
type ConcreteAggregate struct {
	//私有存储容器
	docker []interface{}
}

func (c *ConcreteAggregate) Add(obj interface{}) {
	c.docker = append(c.docker,obj)
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


