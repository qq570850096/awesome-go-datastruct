package main

import "testing"

func TestQueueWrapAround(t *testing.T) {
	queue := NewQueue(3)

	for _, value := range []int{1, 2, 3} {
		if !queue.EnQueue(value) {
			t.Fatalf("EnQueue(%d) failed before queue was full", value)
		}
	}
	if queue.EnQueue(4) {
		t.Fatal("EnQueue should fail when queue is full")
	}

	for _, want := range []int{1, 2} {
		ok, got := queue.DeQueue()
		if !ok || got != want {
			t.Fatalf("DeQueue() = (%v, %d), want (true, %d)", ok, got, want)
		}
	}

	for _, value := range []int{4, 5} {
		if !queue.EnQueue(value) {
			t.Fatalf("EnQueue(%d) failed after wrap-around space opened", value)
		}
	}

	for _, want := range []int{3, 4, 5} {
		ok, got := queue.DeQueue()
		if !ok || got != want {
			t.Fatalf("DeQueue() = (%v, %d), want (true, %d)", ok, got, want)
		}
	}

	if ok, got := queue.DeQueue(); ok || got != 0 {
		t.Fatalf("empty DeQueue() = (%v, %d), want (false, 0)", ok, got)
	}
}
