package BehavioralType

import "fmt"

// 发布订阅，观察者模式
//意思是:定义对象间一种一-对多的依赖关系，使得每当一个对象改变状态，则所有依赖于它的对
//象都会得到通知并被自动更新。
//以生活中的例子来说，就像我们订阅报纸一样，每天有多少人订阅，当有新报纸发布的时候，就
//会有多少人收到新发布的报纸，这种模式就是订阅一发布模式，而报社和订阅者就满足定义中
//说是的，-对多的依赖关系。

// 读者接口（订阅接口）
type IReader interface {
	Update(bookName string)
}

// 读者类（订阅者）
type Reader struct {
	name string
}

func (r *Reader) Update(bookName string) {
	fmt.Println(r.name,"-收到了图书",bookName)
}

// 平台接口（发布方接口）
type IPlatform interface {
	Attach(reader IReader)
	Detach(reader IReader)
	NotifyObservers(bookName string)
}

// 具体发布类（发布方）
type Platform struct {
	list []IReader
}

func (p *Platform) Attach(reader IReader) {
	// 增加读者（订阅者）
	p.list = append(p.list, reader)
}

func (p *Platform) Detach(reader IReader) {
	// 删除读者（订阅者）
	for i,v := range p.list {
		if v == reader {
			// 删除第i个元素,因为interface类型在golang中
			// 以地址的方式传递，所以可以直接比较进行删除
			// golang中只要记得byte,int,bool,string，数组，结构体，默认传值，其他的默认传地址即可
			p.list = append(p.list[:i],p.list[i+1:]...)
		}
	}
}

func (p *Platform) NotifyObservers(bookName string) {
	// 通知所有读者
	for _,reader := range p.list {
		reader.Update(bookName)
	}
}

func (p *Platform) Change (bookName string)  {
	p.NotifyObservers(bookName)
}


