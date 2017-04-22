package chainofresponsibility

import "fmt"

type IHander interface {
	HanderRequest(int)
	SetHander(IHander)
}

type ConcreteHandlerA struct {
	hander IHander
}

func (c *ConcreteHandlerA) SetHander(hander IHander) {
	c.hander = hander
}

func (c *ConcreteHandlerA) HanderRequest(i int) {
	if i < 10 {
		fmt.Println("ConcreteHandlerA handler request", i)
	} else {
		if c.hander != nil {
			c.hander.HanderRequest(i)
		}
	}
}

type ConcreteHandlerB struct {
	hander IHander
}

func (c *ConcreteHandlerB) SetHander(hander IHander) {
	c.hander = hander
}

func (c *ConcreteHandlerB) HanderRequest(i int) {
	if i < 100 {
		fmt.Println("ConcreteHandlerB handler request", i)
	} else {
		if c.hander != nil {
			c.hander.HanderRequest(i)
		}
	}
}

type ConcreteHandlerC struct {
	hander IHander
}

func (c *ConcreteHandlerC) SetHander(hander IHander) {
	c.hander = hander
}

func (c *ConcreteHandlerC) HanderRequest(i int) {
	if i < 1000 {
		fmt.Println("ConcreteHandlerC handler request", i)
	} else {
		if c.hander != nil {
			c.hander.HanderRequest(i)
		}
	}
}
