package StructuralType

import "fmt"

type MobilePhone struct {
	Impl SoftImplementor
}

func (MobilePhone) Run()  {}

type SoftImplementor interface {
	RawRun()
}

type GameSoft struct {
	SoftImplementor
}
type ChatSoft struct {
	SoftImplementor
}

func (GameSoft)RawRun()  {
	fmt.Println("游戏软件启动")
}

func (ChatSoft)RawRun()  {
	fmt.Println("聊天软件启动")
}

type HuaWei struct {
	MobilePhone
}

func (h *HuaWei) Run()  {
	h.Impl.RawRun()
}

func (h *HuaWei) GPUTurbo()  {
	fmt.Println("GPUTurbo started")
	h.Run()
	fmt.Println("GPUTurbo ended")
}

type XiaoMi struct {
	MobilePhone
}

func (x *XiaoMi) Run()  {
	x.Impl.RawRun()
}

func (x *XiaoMi) GameTurbo() {
	fmt.Println("GameTurbo started.")
}

