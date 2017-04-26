package visitor

import "fmt"

type IElement interface {
	Accept(IVisitor)
}

type Employee struct {
	Name         string
	VacationDays int
	Income       float32
}

type Employees struct {
	employees []IElement
}

func (e *Employees) Attach(element IElement) {
	e.employees = append(e.employees, element)
}

func (e *Employees) Detach(element IElement) {
	for k, v := range e.employees {
		if v == element {
			e.employees = append(e.employees[:k], e.employees[k+1:]...)
		}
	}
}

func (e *Employees) Visitor(v IVisitor) {
	for _, el := range e.employees {
		el.Accept(v)
	}
}

func (e *Employee) Accept(v IVisitor) {
	v.Visitor(e)
}

//IVisitor 访问每个元素
type IVisitor interface {
	Visitor(IElement)
}

type IncomeVisitor struct {
}

func (i *IncomeVisitor) Visitor(e IElement) {
	if v, ok := e.(*Employee); ok {
		v.Income = v.Income * 2
		fmt.Printf("%s new income is %f \n", v.Name, v.Income)
	}
}

type VacationVisitor struct {
}

func (v *VacationVisitor) Visitor(e IElement) {
	if v, ok := e.(*Employee); ok {
		v.VacationDays += 5
		fmt.Printf("%s new vacationdays is %d \n", v.Name, v.VacationDays)
	}
}
