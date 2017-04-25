package decorator

//首先定一个行为接口
type IBeverage interface {
	GetDescription() string
	GetPrice() float32
}

//CoffeeBean1 被装饰对象，继承IBeverage行为接口
type CoffeeBean1 struct {
	//string description
}

func (c *CoffeeBean1) GetDescription() string {
	return "first CoffeeBean"
}

func (c *CoffeeBean1) GetPrice() float32 {
	return 30.0
}

//CoffeeBean2 被装饰对象，继承IBeverage行为接口
type CoffeeBean2 struct {
}

func (*CoffeeBean2) GetDescription() string {
	return "second CoffeeBean"
}

func (*CoffeeBean2) GetPrice() float32 {
	return 35.00
}

//Decorator 装饰类
type Decorator struct {
}

func (*Decorator) GetDescription() string {
	return ""
}

func (*Decorator) GetPrice() float32 {
	return 0.0
}

//Milk 具体的牛奶装饰类
type Milk struct {
	Decorator
	beverage IBeverage
}

func (m *Milk) SetBeverage(b IBeverage) IBeverage {
	m.beverage = b
	return m
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + " add milk"
}

func (m *Milk) GetPrice() float32 {
	return m.beverage.GetPrice() + 19.00
}

//Mocha 具体的摩卡装饰类
type Mocha struct {
	Decorator
	beverage IBeverage
}

func (m *Mocha) SetBeverage(b IBeverage) IBeverage {
	m.beverage = b
	return m
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + " add mocha"
}

func (m *Mocha) GetPrice() float32 {
	return m.beverage.GetPrice() + 14.00
}
