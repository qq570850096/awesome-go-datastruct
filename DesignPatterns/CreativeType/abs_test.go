package CreativeType

import (
	"testing"
)

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
