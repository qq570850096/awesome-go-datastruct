package CreativeType

import "testing"
// 抽象工厂的测试类
func TestTCLFactory_NewTV(t *testing.T) {
	TCLfactory := TCLFactory{}
	ref := TCLfactory.NewRefrigerator()
	DoSomething()
	tv := TCLfactory.NewTV()
	DoSomething()

	MEdiaF := MediaFactory{}
	ref = MEdiaF.NewRefrigerator()
	DoSomething()
	tv = MEdiaF.NewTV()
	DoSomething()
}
