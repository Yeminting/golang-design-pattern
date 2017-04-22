package builder

import "fmt"

type Product struct {
	drink     string
	humburger string
}

type IPackageBuilder interface {
	BuildDrink()
	BuildHumburger()
	GetPackage() Product
}

type PackageA struct {
	product Product
}

func (p *PackageA) BuildDrink() {
	p.product.drink = "Coca Cola"
}

func (p *PackageA) BuildHumburger() {
	p.product.humburger = "Chicken Burger"
}

func (p *PackageA) GetPackage() Product {
	return p.product
}

type PackageB struct {
	product Product
}

func (p *PackageB) BuildDrink() {
	p.product.drink = "Pepis Cola"
}

func (p *PackageB) BuildHumburger() {
	p.product.humburger = "Beef Burger"
}

func (p *PackageB) GetPackage() Product {
	return p.product
}

type Director struct {
	builder IPackageBuilder
}

func (d *Director) ConstructPackage(b IPackageBuilder) {
	fmt.Println("begin build package")
	b.BuildDrink()
	b.BuildHumburger()
	d.builder = b
}

func (d *Director) GetPackage() Product {
	return d.builder.GetPackage()
}
