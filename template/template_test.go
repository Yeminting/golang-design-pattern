package template

import "testing"

func TestTemplate(t *testing.T) {
	abstractA := new(ContreteA)
	abstractA.TemplateMethod(abstractA)
	//out
	//ContreteA AbstractMethod1 is call
	//ContreteA AbstractMethod2 is call

	abstractB := new(ContreteB)
	abstractB.TemplateMethod(abstractB)
	//out
	//ContreteB AbstractMethod1 is call
	//ContreteB AbstractMethod2 is call

	//if override TemplateMethod
	//out
	//ContreteB AbstractMethod2 is call
	//ContreteB AbstractMethod1 is call
}
