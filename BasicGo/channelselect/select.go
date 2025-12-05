package channeldemo

import (
	"context"
	"sync"
	"time"
)

// FanIn 把多个输入 channel 合并为一个输出，直到所有输入关闭或 ctx 取消。
// 场景：日志/指标收集器可以把多个微服务的输出汇聚到一个消费者，避免在业务层显式维护众多 select。
func FanIn(ctx context.Context, inputs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range inputs {
		if ch == nil {
			continue
		}
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// TaggedMessage 用在日志聚合场景，标记消息来源。
type TaggedMessage struct {
	Source  string
	Payload string
}

// AggregateLogs 展示 FanIn 的实际应用：把多服务日志合流并标记来源，供后续统一写入 ES/ClickHouse。
func AggregateLogs(ctx context.Context, streams map[string]<-chan string) <-chan TaggedMessage {
	out := make(chan TaggedMessage)
	var wg sync.WaitGroup
	for source, ch := range streams {
		if ch == nil {
			continue
		}
		wg.Add(1)
		go func(name string, c <-chan string) {
			defer wg.Done()
			for msg := range c {
				select {
				case <-ctx.Done():
					return
				case out <- TaggedMessage{Source: name, Payload: msg}:
				}
			}
		}(source, ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Ticker 回传一个只写 channel，展示 chan<- 用法。
// 场景：可以给 WebSocket/长连接发送心跳，或触发周期任务。
func Ticker(ctx context.Context, interval time.Duration) <-chan time.Time {
	ticks := make(chan time.Time)
	go func(ch chan<- time.Time) {
		defer close(ch)
		t := time.NewTicker(interval)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case ts := <-t.C:
				ch <- ts
			}
		}
	}(ticks)
	return ticks
}

// Drain 在 channel 关闭后退出，展示 for-range + close 配合。
func Drain(ctx context.Context, in <-chan int) []int {
	var res []int
	for {
		select {
		case <-ctx.Done():
			return res
		case v, ok := <-in:
			if !ok {
				return res
			}
			res = append(res, v)
		}
	}
}
