package StructuralType

import "fmt"

type MenuComponent interface {
	// 采购设备或者添加子部门
	Add(menuComponent MenuComponent)
	Remove(menuComponent MenuComponent)
	// 查询该节点下所有设备和部门
	GetName() string
	GetPrice() float64
	GetDescription() string
	IsVegetarian() bool
	CreateIterator()
	Display(depth int)
}

type Leaf struct {
	name string
	vegetarian bool
	description string
	price float64
}

func (l *Leaf) Add(menuComponent MenuComponent) {
	panic("Leaf Node can not add")
}

func (l *Leaf) Remove(menuComponent MenuComponent) {
	panic("Leaf Node can not remove")
}

func (l *Leaf) GetName() string {
	return l.name
}

func (l *Leaf) GetPrice() float64 {
	return l.price
}

func (l *Leaf) IsVegetarian() bool {
	return l.vegetarian
}

func (l *Leaf) GetDescription() string {
	return l.description
}

func (l *Leaf) CreateIterator() {
	panic("implement me")
}


func (l *Leaf) Display(depth int) {
	// 输出树形结构的叶子结点，这里直接输出设备名
	for i:=0; i<depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(l.Name())
}

func (l *Leaf) Name() string {
	return l.name
}

func (l *Leaf) SetName(name string) {
	l.name = name
}

// 复合构件
type Composite struct {
	name string
	description string
	arr []MenuComponent
}

func (c *Composite) GetName() string {
	return c.name
}

func (c *Composite) GetPrice() float64 {
	panic("It is not an item.")
}

func (c *Composite) GetDescription() string {
	return c.description
}

func (c *Composite) IsVegetarian() bool {
	panic("implement me")
}

func (c *Composite) CreateIterator() {
	panic("implement me")
}

func (c *Composite) Add(MenuComponent MenuComponent) {
	c.arr = append(c.arr,MenuComponent)
}

func (c *Composite) Remove(MenuComponent MenuComponent) {
	for i,v := range c.arr {
		if v == MenuComponent {
			// 删除第i个元素,因为interface类型在golang中
			// 以地址的方式传递，所以可以直接比较进行删除
			// golang中只要记得byte,int,bool,string，数组，结构体，默认传值，其他的默认传地址即可
			c.arr = append(c.arr[:i],c.arr[i+1:]...)
		}
	}
}

func (c *Composite) Display(depth int) {
	// 输出树形结构
	for i:=0; i<depth; i++ {
		fmt.Print("*")
	}
	fmt.Println(c.GetName())
	// 递归显示
	for _,com := range c.arr {
		com.Display(depth+1)
	}
}

