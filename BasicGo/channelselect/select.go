package channeldemo

import (
	"context"
	"sync"
	"time"
)

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

type TaggedMessage struct {
	Source  string
	Payload string
}

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
