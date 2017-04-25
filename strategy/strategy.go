package strategy

import "fmt"

type IStrategy interface {
	Action()
}

type Strategy1 struct {
}

func (*Strategy1) Action() {
	fmt.Println("strategy1 call")
}

type Strategy2 struct {
}

func (*Strategy2) Action() {
	fmt.Println("strategy2 call")
}

type Context struct {
	strategy IStrategy
}

func (c *Context) SetStrategy(s IStrategy) {
	c.strategy = s
}

func (c *Context) CallStrategy() {
	c.strategy.Action()
}
