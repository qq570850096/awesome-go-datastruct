package BehavioralType

import "fmt"

type IReader interface {
	Update(bookName string)
}

type Reader struct {
	name string
}

func (r *Reader) Update(bookName string) {
	fmt.Println(r.name, " received book ", bookName)
}

type IPlatform interface {
	Attach(reader IReader)
	Detach(reader IReader)
	NotifyObservers(bookName string)
}

type Platform struct {
	list []IReader
}

func (p *Platform) Attach(reader IReader) {

	p.list = append(p.list, reader)
}

func (p *Platform) Detach(reader IReader) {

	for i, v := range p.list {
		if v == reader {

			p.list = append(p.list[:i], p.list[i+1:]...)
		}
	}
}

func (p *Platform) NotifyObservers(bookName string) {

	for _, reader := range p.list {
		reader.Update(bookName)
	}
}

func (p *Platform) Change(bookName string) {
	p.NotifyObservers(bookName)
}
