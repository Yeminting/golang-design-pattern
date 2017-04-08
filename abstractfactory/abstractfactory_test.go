package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	var (
		factory IFactory
		product IProduct
	)

	factory = new(FactoryA)
	product = factory.CreateProduct()
	product.ShowProduct()
	//out
	//i am product a

	factory = new(FactoryB)
	product = factory.CreateProduct()
	product.ShowProduct()
	//out
	//i am product b
}
