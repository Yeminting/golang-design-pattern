package template

import "fmt"

type IAbstract interface {
	AbstractMethod1()
	AbstractMethod2()
}

type AbstractClass struct {
}

func (*AbstractClass) AbstractMethod1() {

}
func (*AbstractClass) AbstractMethod2() {

}

func (a *AbstractClass) TemplateMethod(c IAbstract) {
	c.AbstractMethod1()
	c.AbstractMethod2()
}

type ContreteA struct {
	AbstractClass
}

func (*ContreteA) AbstractMethod1() {
	fmt.Println("ContreteA AbstractMethod1 is call")
}

func (*ContreteA) AbstractMethod2() {
	fmt.Println("ContreteA AbstractMethod2 is call")
}

type ContreteB struct {
	AbstractClass
}

func (*ContreteB) AbstractMethod1() {
	fmt.Println("ContreteB AbstractMethod1 is call")
}

func (*ContreteB) AbstractMethod2() {
	fmt.Println("ContreteB AbstractMethod2 is call")
}

//TemplateMethod if override TemplateMethod what happend.
func (b *ContreteB) TemplateMethod(a IAbstract) {
	b.AbstractMethod2()
	b.AbstractMethod1()
}
