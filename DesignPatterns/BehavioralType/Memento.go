package BehavioralType

// 备忘录
type Memento struct {
	state string // 这里就是保存的状态
}

func (m *Memento) SetState(s string) {
	m.state = s
}

func (m *Memento) GetState() string {
	return m.state
}

// 发起人
type Originator struct {
	state string // 这里就简单一点，要保存的状态就是一个字符串
}

func (o *Originator) SetState(s string) {
	o.state = s
}

func (o *Originator) GetState() string {
	return o.state
}

// 这里就是规定了要保存的状态范围
func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

// 负责人
type Caretaker struct {
	memento *Memento
}

func (c *Caretaker) GetMemento() *Memento {
	return c.memento
}

func (c *Caretaker) SetMemento(m *Memento) {
	c.memento = m
}