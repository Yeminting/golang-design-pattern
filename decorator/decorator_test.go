package decorator

import "testing"

func TestDecorator(t *testing.T) {
	var berevage IBeverage
	berevage = new(CoffeeBean1)
	berevage = new(Milk).SetBeverage(berevage)
	t.Log("\n", berevage.GetDescription(), "价格", berevage.GetPrice())
	//out
	//first CoffeeBean add milk 价格 49

	berevage = new(Mocha).SetBeverage(berevage)
	t.Log("\n", berevage.GetDescription(), "价格", berevage.GetPrice())
	//out
	//first CoffeeBean add milk add mocha 价格 63
}
