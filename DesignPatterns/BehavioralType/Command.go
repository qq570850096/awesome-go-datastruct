package BehavioralType

import "fmt"

type command interface {
	Execute()
}
type TV struct {}

func (tv *TV) Open()  {
	fmt.Println("打开电视")
}
func (tv *TV) Close()  {
	fmt.Println("关闭电视")
}
func (tv *TV)Change()  {
	fmt.Println("换台")
}

// 接受者
type OpenTvCommand struct {
	tv *TV
}

func (o *OpenTvCommand) Execute()  {
	o.tv.Open()
}
type CloseTvCommand struct {
	tv *TV
}

func (c *CloseTvCommand) Execute()  {
	c.tv.Close()
}

type ChangeTvCommand struct {
	tv *TV
}

func (c *ChangeTvCommand) Execute()  {
	c.tv.Change()
}

type TVRemote struct {
	open *OpenTvCommand
	change *ChangeTvCommand
	close *CloseTvCommand
}

func (tv *TVRemote) Open ()  {
	tv.open.Execute()
}
func (tv *TVRemote) Change()  {
	tv.change.Execute()
}
func (tv *TVRemote) Close ()  {
	tv.close.Execute()
}



