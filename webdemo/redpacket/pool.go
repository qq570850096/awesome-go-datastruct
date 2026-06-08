package redpacket

import (
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Pool struct {
	mu sync.Mutex

	remainingAmount int64
	remainingCount  int
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Init(totalAmount int64, count int) error {
	if totalAmount <= 0 || count <= 0 {
		return errors.New("totalAmount and count must be positive")
	}
	if totalAmount < int64(count) {
		return errors.New("totalAmount must be at least count (1 unit each)")
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.remainingAmount = totalAmount
	p.remainingCount = count
	return nil
}

func (p *Pool) Grab() (int64, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.remainingCount == 0 {
		return 0, errors.New("no red packets left")
	}

	if p.remainingCount == 1 {
		amount := p.remainingAmount
		p.remainingAmount = 0
		p.remainingCount = 0
		return amount, nil
	}

	max := p.remainingAmount / int64(p.remainingCount) * 2
	if max <= 1 {
		max = 1
	}
	amount := rand.Int63n(max-1) + 1

	if p.remainingAmount-amount < int64(p.remainingCount-1) {
		amount = p.remainingAmount - int64(p.remainingCount-1)
	}

	p.remainingAmount -= amount
	p.remainingCount--
	return amount, nil
}

func (p *Pool) Stats() (amount int64, count int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.remainingAmount, p.remainingCount
}

type PoolV2 struct {
	amounts []int64
	index   int64
}

func (p *PoolV2) Init(totalAmount int64, count int) error {
	if totalAmount <= 0 || count <= 0 {
		return errors.New("totalAmount and count must be positive")
	}
	if totalAmount < int64(count) {
		return errors.New("totalAmount must be at least count (1 unit each)")
	}
	p.amounts = make([]int64, count)

	remainAmount := totalAmount
	remainCount := count
	for i := 0; i < count; i++ {
		if remainCount == 1 {
			p.amounts[i] = remainAmount
			break
		}

		max := remainAmount / int64(remainCount) * 2
		if max <= 1 {
			max = 1
		}
		amount := rand.Int63n(max-1) + 1

		if remainAmount-amount < int64(remainCount-1) {
			amount = remainAmount - int64(remainCount-1)
		}

		p.amounts[i] = amount
		remainAmount -= amount
		remainCount--
	}
	atomic.StoreInt64(&p.index, 0)
	return nil
}

func (p *PoolV2) Grab() (int64, error) {
	i := atomic.AddInt64(&p.index, 1) - 1
	if i < 0 || int(i) >= len(p.amounts) {
		return 0, errors.New("no red packets left")
	}
	return p.amounts[i], nil
}
