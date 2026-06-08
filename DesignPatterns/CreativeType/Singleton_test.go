package CreativeType

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance = GetInstance()
	instance.IsEmpty()
	instance.IsBoiled()

	instance.Drain()

	instance.Fill()
	instance.Boil()

	instance = GetInstance()
	instance.Drain()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			instance = GetInstance()
			fmt.Println(instance)
			wg.Done()
		}()
	}
	wg.Wait()
}
