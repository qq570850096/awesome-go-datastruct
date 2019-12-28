package CreativeType

import (
	"testing"
)
// 抽象工厂的测试类
func TestTCLFactory_NewTV(t *testing.T) {
	var factory Factory
	factory = &TCLFactory{}
	ref := factory.NewRefrigerator()
	ref.DoSomething()
	tv := factory.NewTV()
	tv.DoSomething()
	factory = &MediaFactory{}
	ref = factory.NewRefrigerator()
	ref.DoSomething()
	tv = factory.NewTV()
	tv.DoSomething()

}
