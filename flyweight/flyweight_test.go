package flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	factory := FlyWeightFactory{}
	f := factory.GetFlyWeight("A")
	f.Operation()
	f = factory.GetFlyWeight("B")
	f.Operation()

	f = factory.GetFlyWeight("A")
	f.Operation()

	//out
	//A object is call
	//B object is call
	//A object is call
}
