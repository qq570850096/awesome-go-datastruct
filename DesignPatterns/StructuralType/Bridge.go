package StructuralType

import "fmt"

// MobilePhone is the abstraction side of the bridge. It keeps a reference to a
// software implementor so phone brands and software can vary independently.
type MobilePhone struct {
	Impl SoftImplementor
}

func (MobilePhone) Run() {}

// SoftImplementor is the implementation side of the bridge.
type SoftImplementor interface {
	RawRun()
}

// GameSoft and ChatSoft are concrete software implementors.
type GameSoft struct {
	SoftImplementor
}
type ChatSoft struct {
	SoftImplementor
}

func (GameSoft) RawRun() {
	fmt.Println("game app started")
}

func (ChatSoft) RawRun() {
	fmt.Println("chat app started")
}

type HuaWei struct {
	MobilePhone
}

// Run delegates to the bridged software implementation.
func (h *HuaWei) Run() {
	h.Impl.RawRun()
}

func (h *HuaWei) GPUTurbo() {
	fmt.Println("GPUTurbo started")
	h.Run()
	fmt.Println("GPUTurbo ended")
}

type XiaoMi struct {
	MobilePhone
}

func (x *XiaoMi) Run() {
	x.Impl.RawRun()
}

func (x *XiaoMi) GameTurbo() {
	fmt.Println("GameTurbo started.")
}
