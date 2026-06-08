package CreativeType

import "testing"

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
