package CreativeType

import "testing"
// 抽象工厂的测试类
func TestTCLFactory_NewTV(t *testing.T) {
	TCLfactory := TCLFactory{}
	ref := TCLfactory.NewRefrigerator()
	ref.DoSomething()
	tv := TCLfactory.NewTV()
	tv.DoSomething()

	MEdiaF := MediaFactory{}
	ref = MEdiaF.NewRefrigerator()
	ref.DoSomething()
	tv = MEdiaF.NewTV()
	tv.DoSomething()
}
