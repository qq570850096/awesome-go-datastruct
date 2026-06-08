package BehavioralType

import (
	"fmt"
	"reflect"
)

type MeditorCompany interface {
	GetSeller() Colleaguer
	SetSeller(seller ColleagueSeller)
	GetBuyer() Colleaguer
	SetBuyer(ColleagueBuyer)
	GetName() string
	SetName(name string)
	Publish(message string, colleaguer Colleaguer)
}

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
