package CreativeType

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance = GetInstance()
	instance.IsEmpty()
	instance.IsBoiled()

	instance.Drain()

	instance.Fill()
	instance.Boil()
	// 这里再次get会发现返回的依然是我们的instance
	instance = GetInstance()
	instance.Drain()
}

