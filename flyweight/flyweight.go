package flyweight

import "fmt"

type IFlyWeight interface {
	Operation()
}

type ConcreteFlyWeight struct {
	state string
}

func (c *ConcreteFlyWeight) Operation() {
	fmt.Printf("%s object is call \n", c.state)
}

type FlyWeightFactory struct {
	//object pool
	flyweights map[string]IFlyWeight
}

func (f *FlyWeightFactory) GetFlyWeight(s string) IFlyWeight {
	if f.flyweights == nil {
		f.flyweights = make(map[string]IFlyWeight)
	}

	//if exist return
	if v, ok := f.flyweights[s]; ok {
		return v
	}

	//if not exist then new
	fly := new(ConcreteFlyWeight)
	fly.state = s
	f.flyweights[s] = fly

	return fly
}
