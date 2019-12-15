package StructuralType

import "testing"

func TestMobilePhone_Run(t *testing.T) {
	h1 := HuaWei{MobilePhone{ChatSoft{}}}
	h2 := HuaWei{MobilePhone{GameSoft{}}}
	h1.Run()
	h2.Run()
	h2.GPUTurbo()
	m1 := XiaoMi{MobilePhone{ChatSoft{}}}
	m2 := XiaoMi{MobilePhone{GameSoft{}}}
	m1.Run()
	m2.Run()
	m2.GameTurbo()
}
