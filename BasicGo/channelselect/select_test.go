package channeldemo

import (
	"context"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	makeStream := func(values ...string) <-chan string {
		ch := make(chan string, len(values))
		for _, v := range values {
			ch <- v
		}
		close(ch)
		return ch
	}

	out := FanIn(ctx, makeStream("a", "b"), makeStream("c"))
	got := make(map[string]bool)
	for v := range out {
		got[v] = true
	}
	if len(got) != 3 || !got["a"] || !got["b"] || !got["c"] {
		t.Fatalf("unexpected fan-in result: %#v", got)
	}
}

func TestTickerAndDrain(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()
	ticks := Ticker(ctx, time.Microsecond)
	count := 0
	for range ticks {
		count++
		if count > 2 {
			break
		}
	}
	if count == 0 {
		t.Fatalf("expected at least one tick, got %d", count)
	}

	ints := make(chan int, 3)
	ints <- 1
	ints <- 2
	close(ints)
	values := Drain(context.Background(), ints)
	if len(values) != 2 || values[0] != 1 || values[1] != 2 {
		t.Fatalf("unexpected drain result %#v", values)
	}
}

func TestAggregateLogs(t *testing.T) {
	ctx := context.Background()
	streamA := make(chan string, 2)
	streamB := make(chan string, 1)
	streamA <- "a1"
	streamA <- "a2"
	streamB <- "b1"
	close(streamA)
	close(streamB)

	streams := map[string]<-chan string{
		"svcA": streamA,
		"svcB": streamB,
	}
	out := AggregateLogs(ctx, streams)
	collected := map[string][]string{}
	for msg := range out {
		collected[msg.Source] = append(collected[msg.Source], msg.Payload)
	}
	if len(collected["svcA"]) != 2 || collected["svcA"][0] != "a1" || collected["svcA"][1] != "a2" {
		t.Fatalf("unexpected svcA logs %#v", collected["svcA"])
	}
	if len(collected["svcB"]) != 1 || collected["svcB"][0] != "b1" {
		t.Fatalf("unexpected svcB logs %#v", collected["svcB"])
	}
}
