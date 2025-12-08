package redpacket

import (
	"errors"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 初始化全局随机数种子，避免在高并发路径上重复调用 rand.Seed。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Pool 表示一个简单的红包池，使用整数金额（分）避免浮点误差。
type Pool struct {
	mu sync.Mutex

	remainingAmount int64
	remainingCount  int
}

// NewPool 创建一个空的红包池。
func NewPool() *Pool {
	return &Pool{}
}

// Init 初始化红包池，总金额（分）和红包个数。
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

// Grab 尝试抢一个红包，返回金额（分）。
// 使用一个很简单的“二倍均值”随机算法。
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

	// 二倍均值算法：随机 [1, 2*avg-1]
	max := p.remainingAmount / int64(p.remainingCount) * 2
	if max <= 1 {
		max = 1
	}
	amount := rand.Int63n(max-1) + 1

	// 保证至少给剩余红包每个 1 单位
	if p.remainingAmount-amount < int64(p.remainingCount-1) {
		amount = p.remainingAmount - int64(p.remainingCount-1)
	}

	p.remainingAmount -= amount
	p.remainingCount--
	return amount, nil
}

// Stats 返回剩余金额和剩余个数，用于监控和测试。
func (p *Pool) Stats() (amount int64, count int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.remainingAmount, p.remainingCount
}

// PoolV2 通过预分配所有红包金额，并使用原子操作分发，避免在 Grab 阶段加锁。
type PoolV2 struct {
	amounts []int64
	index   int64
}

// Init 按二倍均值算法预先生成所有红包金额。
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

// Grab 使用原子自增索引分发预生成的红包金额。
func (p *PoolV2) Grab() (int64, error) {
	i := atomic.AddInt64(&p.index, 1) - 1
	if i < 0 || int(i) >= len(p.amounts) {
		return 0, errors.New("no red packets left")
	}
	return p.amounts[i], nil
}

