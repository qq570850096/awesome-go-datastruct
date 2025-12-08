package redpacket

import (
	"sync"
	"testing"
)

func TestPoolGrabConcurrent(t *testing.T) {
	const totalAmount = int64(1000)
	const count = 100

	p := NewPool()
	if err := p.Init(totalAmount, count); err != nil {
		t.Fatalf("Init error: %v", err)
	}

	var wg sync.WaitGroup
	result := make([]int64, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			amount, err := p.Grab()
			if err != nil {
				// 允许极端情况下有并发失败，但不影响总和校验
				return
			}
			result[idx] = amount
		}(i)
	}
	wg.Wait()

	var sum int64
	var nonZero int
	for _, v := range result {
		sum += v
		if v > 0 {
			nonZero++
		}
	}

	if sum != totalAmount {
		t.Fatalf("sum = %d, want %d", sum, totalAmount)
	}
	if nonZero != count {
		t.Fatalf("expected %d successful grabs, got %d", count, nonZero)
	}

	remainAmount, remainCount := p.Stats()
	if remainAmount != 0 || remainCount != 0 {
		t.Fatalf("pool should be empty, got amount=%d count=%d", remainAmount, remainCount)
	}
}

func TestPoolV2GrabConcurrent(t *testing.T) {
	const totalAmount = int64(1000)
	const count = 100

	var p PoolV2
	if err := p.Init(totalAmount, count); err != nil {
		t.Fatalf("Init error: %v", err)
	}

	var wg sync.WaitGroup
	result := make([]int64, count)

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			amount, err := p.Grab()
			if err != nil {
				return
			}
			result[idx] = amount
		}(i)
	}
	wg.Wait()

	var sum int64
	var nonZero int
	for _, v := range result {
		sum += v
		if v > 0 {
			nonZero++
		}
	}
	if sum != totalAmount {
		t.Fatalf("sum = %d, want %d", sum, totalAmount)
	}
	if nonZero != count {
		t.Fatalf("expected %d successful grabs, got %d", count, nonZero)
	}
}

func BenchmarkPoolGrab(b *testing.B) {
	const totalAmount = int64(1000000)
	const count = 100000

	for i := 0; i < b.N; i++ {
		p := NewPool()
		if err := p.Init(totalAmount, count); err != nil {
			b.Fatalf("Init error: %v", err)
		}
		for j := 0; j < count; j++ {
			if _, err := p.Grab(); err != nil {
				break
			}
		}
	}
}

func BenchmarkPoolV2Grab(b *testing.B) {
	const totalAmount = int64(1000000)
	const count = 100000

	for i := 0; i < b.N; i++ {
		var p PoolV2
		if err := p.Init(totalAmount, count); err != nil {
			b.Fatalf("Init error: %v", err)
		}
		for j := 0; j < count; j++ {
			if _, err := p.Grab(); err != nil {
				break
			}
		}
	}
}

func BenchmarkPoolGrabParallel(b *testing.B) {
	const totalAmount = int64(1000000)
	const count = 100000

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p := NewPool()
			if err := p.Init(totalAmount, count); err != nil {
				b.Fatalf("Init error: %v", err)
			}
			var wg sync.WaitGroup
			for i := 0; i < count; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					_, _ = p.Grab()
				}()
			}
			wg.Wait()
		}
	})
}

