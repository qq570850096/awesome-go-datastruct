package BehavioralType

import "fmt"

// IReader is the observer role. It receives updates from the platform.
type IReader interface {
	Update(bookName string)
}

// Reader is a concrete observer.
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

// Platform is the subject role. It stores observers and broadcasts changes.
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

// NotifyObservers pushes the new book name to every registered reader.
func (p *Platform) NotifyObservers(bookName string) {

	for _, reader := range p.list {
		reader.Update(bookName)
	}
}

// Change is the subject's state-change entrypoint.
func (p *Platform) Change(bookName string) {
	p.NotifyObservers(bookName)
}
