package abstractfactory

import "fmt"

type IProduct interface {
	ShowProduct()
}

type ProductA struct {
}

func (*ProductA) ShowProduct() {
	fmt.Println("i am product a")
}

type ProductB struct {
}

func (*ProductB) ShowProduct() {
	fmt.Println("i am product b")
}

//IFactory is abstract factory.
type IFactory interface {
	CreateProduct() IProduct
}

//FactoryA create product a.
type FactoryA struct {
}

func (*FactoryA) CreateProduct() IProduct {
	return new(ProductA)
}

//FactoryB create product b.
type FactoryB struct {
}

func (*FactoryB) CreateProduct() IProduct {
	return new(ProductB)
}
