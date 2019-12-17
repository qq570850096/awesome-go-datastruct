package CreativeType

import "testing"
// 因为go语言本身支持指针操作，对于clone的操作还是很方便的
func TestRoleChinese_Clone(t *testing.T) {
	role := &RoleChinese{
		HeadColor: "black",
		EyesColor: "black",
		Tall:      170,
	}
	role.Show()
	copyer := role.Clone()
	copyer.Show()
	copyer.SetEyesColor("bule")
	role.Show()
	copyer.Show()
}
