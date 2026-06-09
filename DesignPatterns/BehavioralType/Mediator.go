package BehavioralType

import (
	"fmt"
	"reflect"
)

// MeditorCompany is the mediator interface. Colleagues talk to the mediator
// instead of calling each other directly.
type MeditorCompany interface {
	GetSeller() Colleaguer
	SetSeller(seller ColleagueSeller)
	GetBuyer() Colleaguer
	SetBuyer(ColleagueBuyer)
	GetName() string
	SetName(name string)
	Publish(message string, colleaguer Colleaguer)
}

// Meditor coordinates buyer and seller message delivery.
type Meditor struct {
	name   string
	buyer  *ColleagueBuyer
	seller *ColleagueSeller
}

func (m *Meditor) SetSeller(seller ColleagueSeller) {
	m.seller = &seller
}

func (m *Meditor) SetBuyer(b ColleagueBuyer) {
	m.buyer = &b
}

// Publish routes a message from one colleague to the opposite colleague.
func (m *Meditor) Publish(message string, colleaguer Colleaguer) {

	if reflect.DeepEqual(colleaguer, m.seller) {
		m.buyer.Accept(message)
	} else if reflect.DeepEqual(colleaguer, m.buyer) {
		m.seller.Accept(message)
	}
}

func (m *Meditor) GetSeller() Colleaguer {
	return m.seller
}

func (m *Meditor) GetBuyer() Colleaguer {
	return m.buyer
}

func (m *Meditor) GetName() string {
	return m.name
}

func (m *Meditor) SetName(name string) {
	m.name = name
}

type Colleaguer interface {
	Colleguer(meditor MeditorCompany)
	Send(string)
	Accept(string)
}

// ColleagueSeller sends and receives messages through a mediator.
type ColleagueSeller struct {
	meditor MeditorCompany
}

func (c *ColleagueSeller) Send(message string) {
	c.meditor.Publish(message, c)
}

func (c *ColleagueSeller) Accept(message string) {
	fmt.Println("seller received message: " + message)
}

func (c *ColleagueSeller) Colleguer(meditor MeditorCompany) {
	c.meditor = meditor
}

type ColleagueBuyer struct {
	meditor MeditorCompany
}

func (c *ColleagueBuyer) Colleguer(meditor MeditorCompany) {
	c.meditor = meditor
}

func (c *ColleagueBuyer) Send(message string) {
	c.meditor.Publish(message, c)
}

func (c *ColleagueBuyer) Accept(message string) {
	fmt.Println("buyer received message: " + message)
}
