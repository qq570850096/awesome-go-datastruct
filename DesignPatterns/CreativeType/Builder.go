package CreativeType

import "fmt"

// Product is the complex object assembled by the builder steps.
type Product struct {
	ground string
	cement string
	roof   string
}

func (p *Product) Cement() string {
	return p.cement
}

func (p *Product) SetCement(cement string) {
	p.cement = cement
}

func (p *Product) Roof() string {
	return p.roof
}

func (p *Product) SetRoof(roof string) {
	p.roof = roof
}

func (p *Product) Ground() string {
	return p.ground
}

func (p *Product) SetGround(ground string) {
	p.ground = ground
}

type Builder interface {
	BuildGround()
	BuildCement()
	BuildRoof()
	BuildProduct() *Product
}

// ConcreteBuilder stores the product under construction and implements each
// construction step in the Builder interface.
type ConcreteBuilder struct {
	p *Product
}

func (this *ConcreteBuilder) BuildGround() {
	this.p.SetGround("build foundation")
	fmt.Println(this.p.ground)
}

func (this *ConcreteBuilder) BuildCement() {
	this.p.SetCement("build house")
	fmt.Println(this.p.Cement())
}
func (this *ConcreteBuilder) BuildRoof() {
	this.p.SetRoof("build roof")
	fmt.Println(this.p.Roof())
}

func (this *ConcreteBuilder) BuildProduct() *Product {
	fmt.Println("build complete")
	return this.p
}

type Director struct {
	builder Builder
}

// Construst coordinates the build order. The misspelled name is kept so the
// existing tests and examples remain stable.
func (this *Director) Construst() Product {
	this.builder.BuildGround()
	this.builder.BuildCement()
	this.builder.BuildRoof()

	return *this.builder.BuildProduct()
}
