package CreativeType

import "testing"
// 简单工厂的测试类
func TestGirlFactory_CreateGirl(t *testing.T) {
	factor := &GirlFactory{}

	Fat := factor.CreateGirl("fat")
	Fat.weight()
	Thin := factor.CreateGirl("thin")
	Thin.weight()
}
