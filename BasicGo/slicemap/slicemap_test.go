package slicemap

import "testing"

func TestMakeNilAndEmpty(t *testing.T) {
	nilSlice, emptySlice := MakeNilAndEmpty()
	if nilSlice != nil {
		t.Fatalf("nilSlice should be nil")
	}
	if emptySlice == nil {
		t.Fatalf("emptySlice should not be nil")
	}
	if len(nilSlice) != 0 || len(emptySlice) != 0 {
		t.Fatalf("both slices should have len 0")
	}
}

func TestShareUnderlying(t *testing.T) {
	base, sub, grown := ShareUnderlying()
	if base[0] != 10 || sub[0] != 10 {
		t.Fatalf("modifying sub should affect base, got base[0]=%d sub[0]=%d", base[0], sub[0])
	}
	if len(grown) < 3 {
		t.Fatalf("grown slice length unexpected: %d", len(grown))
	}
}

func TestFilterInPlace(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	res := FilterInPlace(nums, func(v int) bool { return v%2 == 1 })
	if len(res) != 3 || res[0] != 1 || res[1] != 3 || res[2] != 5 {
		t.Fatalf("FilterInPlace result unexpected: %#v", res)
	}
}

func TestCountWordsAndSet(t *testing.T) {
	text := "go go is fun"
	m := CountWords(text)
	if m["go"] != 2 || m["is"] != 1 || m["fun"] != 1 {
		t.Fatalf("unexpected CountWords map: %#v", m)
	}

	set := NewSet()
	set.Add("a")
	set.Add("b")
	if !set.Has("a") || !set.Has("b") || set.Has("c") {
		t.Fatalf("Set basic operations failed")
	}
	set.Remove("a")
	if set.Has("a") {
		t.Fatalf("Remove did not delete element")
	}
}

