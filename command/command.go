package command

import "fmt"

type ICommand interface {
	Excute()
	Cancel()
}

type Revicer struct {
}

func (r *Revicer) Do() {
	fmt.Println("revicer do something")
}

func (r *Revicer) UnDo() {
	fmt.Println("revicer undo something")
}

type CreateCommand struct {
	revicer *Revicer
}

func (c *CreateCommand) SerRevicer(s *Revicer) {
	c.revicer = s
}

func (c *CreateCommand) Excute() {
	c.revicer.Do()
}

func (c *CreateCommand) Cancel() {
	c.revicer.UnDo()
}

type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCmd(c ICommand) {
	i.command = c
}

func (i *Invoker) ExcuteCmd() {
	i.command.Excute()
}

func (i *Invoker) CancelCmd() {
	i.command.Cancel()
}
