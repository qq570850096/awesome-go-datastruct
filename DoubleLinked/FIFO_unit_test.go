package DoubleLinked

import "testing"

func TestFIFOEvictsOldestEntry(t *testing.T) {
	cache := InitFIFO(2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("c", 3)

	if _, ok := cache.find["a"]; ok {
		t.Fatal("expected oldest key a to be evicted")
	}
	if got := cache.Get("b"); got != 2 {
		t.Fatalf("expected key b to remain, got %v", got)
	}
	if got := cache.Get("c"); got != 3 {
		t.Fatalf("expected key c to remain, got %v", got)
	}
}

func TestFIFODuplicatePutKeepsOrder(t *testing.T) {
	cache := InitFIFO(2)
	cache.Put("a", 1)
	cache.Put("b", 2)
	cache.Put("a", 10)
	cache.Put("c", 3)

	if _, ok := cache.find["a"]; ok {
		t.Fatal("expected updated oldest key a to be evicted")
	}
	if got := cache.Get("b"); got != 2 {
		t.Fatalf("expected key b to remain, got %v", got)
	}
}
