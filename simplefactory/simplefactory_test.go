package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	f := Factory{}
	producta := f.CreateProduct("A")
	producta.ShowProduct()
	//out
	// i am product a

	productb := f.CreateProduct("B")
	productb.ShowProduct()
	//out
	//i am product b
}
