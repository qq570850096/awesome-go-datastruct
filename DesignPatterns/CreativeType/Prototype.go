package CreativeType

import "fmt"

type Role interface {
	Clone() Role
	SetHeadColor(string)
	SetEyesColor(string)
	SetTall(int64)
	Show()
}

type RoleChinese struct {
	HeadColor string
	EyesColor string
	Tall      int64
}

func (pR *RoleChinese) Clone() Role {
	var pChinese = &RoleChinese{HeadColor: pR.HeadColor, EyesColor: pR.EyesColor, Tall: pR.Tall}
	return pChinese
}

func (pR *RoleChinese) SetHeadColor(color string) {
	pR.HeadColor = color
}

func (pR *RoleChinese) SetEyesColor(color string) {
	pR.EyesColor = color
}

func (pR *RoleChinese) SetTall(tall int64) {
	pR.Tall = tall
}

func (pR *RoleChinese) Show() {
	fmt.Println("Role's headcolor is:", pR.HeadColor, " EyesColor is:", pR.EyesColor, " tall is:", pR.Tall)
}
