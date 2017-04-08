package simplefactory

import "fmt"

//IProduct is abstract for product
type IProduct interface {
	ShowProduct()
}

type ProductA struct {
}

func (p *ProductA) ShowProduct() {
	fmt.Println("i am product a")
}

type ProductB struct {
}

func (p *ProductB) ShowProduct() {
	fmt.Println("i am product b")
}

//Factory is static simple factory
type Factory struct {
}

func (f *Factory) CreateProduct(product string) IProduct {
	switch product {
	case "A":
		return new(ProductA)
	case "B":
		return new(ProductB)
	default:
		return nil
	}
}
