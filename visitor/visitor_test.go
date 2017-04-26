package visitor

import "testing"

func TestVisitor(t *testing.T) {
	e := Employees{}
	e.Attach(&Employee{Name: "Hank", VacationDays: 10, Income: 10000})
	e.Attach(&Employee{Name: "Lucy", VacationDays: 5, Income: 8900})
	e.Attach(&Employee{Name: "Lily", VacationDays: 13, Income: 14000})

	v1 := new(IncomeVisitor)
	v2 := new(VacationVisitor)

	e.Visitor(v1)
	//out
	//Hank new income is 20000.000000
	//Lucy new income is 17800.000000
	//Lily new income is 28000.000000

	e.Visitor(v2)
	//out
	//Hank new vacationdays is 15
	//Lucy new vacationdays is 10
	//Lily new vacationdays is 18
}
