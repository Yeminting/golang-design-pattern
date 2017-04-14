package state

import "fmt"

type IState interface {
	HandleState(string)
}

type ConcreteStateBegin struct {
}

func (*ConcreteStateBegin) HandleState(s string) {
	fmt.Println("begin state", s)
}

type ConcreteStateLoading struct {
}

func (*ConcreteStateLoading) HandleState(s string) {
	fmt.Println("loading state", s)
}

type ConcreteStateEnd struct {
}

func (*ConcreteStateEnd) HandleState(s string) {
	fmt.Println("end state", s)
}

type Context struct {
	state IState
}

func (c *Context) ChangeState(s IState) {
	c.state = s
}

func (c *Context) Request(s string) {
	if c.state != nil {
		c.state.HandleState(s)
	}
}
