package sharedvars

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestCounterIsSafeForConcurrentUse(t *testing.T) {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()

	if got := counter.Value(); got != 100 {
		t.Fatalf("Counter.Value() = %d, want 100", got)
	}
}

func TestAtomicCounter(t *testing.T) {
	var counter AtomicCounter
	for i := 0; i < 5; i++ {
		counter.Inc()
	}
	if got := counter.Value(); got != 5 {
		t.Fatalf("AtomicCounter.Value() = %d, want 5", got)
	}
}

func TestOnceValueInitializesOnce(t *testing.T) {
	var onceValue OnceValue
	var calls atomic.Int64

	first := onceValue.Init(func() string {
		calls.Add(1)
		return "first"
	})
	second := onceValue.Init(func() string {
		calls.Add(1)
		return "second"
	})

	if first != "first" || second != "first" {
		t.Fatalf("Init results = (%q, %q), want first twice", first, second)
	}
	if got := calls.Load(); got != 1 {
		t.Fatalf("initializer calls = %d, want 1", got)
	}
}
