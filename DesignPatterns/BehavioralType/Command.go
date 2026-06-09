package BehavioralType

import "fmt"

// command is the command role. It hides the concrete receiver action behind a
// single Execute method.
type command interface {
	Execute()
}

// TV is the receiver role. Concrete commands delegate real work to it.
type TV struct{}

func (tv *TV) Open() {
	fmt.Println("turn on TV")
}
func (tv *TV) Close() {
	fmt.Println("turn off TV")
}
func (tv *TV) Change() {
	fmt.Println("change channel")
}

type OpenTvCommand struct {
	tv *TV
}

// Execute binds the "open TV" request to the receiver action.
func (o *OpenTvCommand) Execute() {
	o.tv.Open()
}

type CloseTvCommand struct {
	tv *TV
}

// Execute binds the "close TV" request to the receiver action.
func (c *CloseTvCommand) Execute() {
	c.tv.Close()
}

type ChangeTvCommand struct {
	tv *TV
}

// Execute binds the "change channel" request to the receiver action.
func (c *ChangeTvCommand) Execute() {
	c.tv.Change()
}

// TVRemote is the invoker role. It exposes high-level buttons and knows only
// about commands, not the TV implementation details.
type TVRemote struct {
	open   *OpenTvCommand
	change *ChangeTvCommand
	close  *CloseTvCommand
}

func (tv *TVRemote) Open() {
	tv.open.Execute()
}
func (tv *TVRemote) Change() {
	tv.change.Execute()
}
func (tv *TVRemote) Close() {
	tv.close.Execute()
}
