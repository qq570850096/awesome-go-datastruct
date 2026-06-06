package pointers

import (
	"slices"
	"testing"
)

func TestIncrementValueDoesNotMutateCaller(t *testing.T) {
	n := 10
	if got := IncrementValue(n); got != 11 {
		t.Fatalf("IncrementValue() = %d, want 11", got)
	}
	if n != 10 {
		t.Fatalf("value parameter should not mutate caller, got %d", n)
	}
}

func TestIncrementPointerMutatesCaller(t *testing.T) {
	n := 10
	if !IncrementPointer(&n) {
		t.Fatal("IncrementPointer should accept non-nil pointer")
	}
	if n != 11 {
		t.Fatalf("pointer parameter should mutate caller, got %d", n)
	}
	if IncrementPointer(nil) {
		t.Fatal("IncrementPointer should reject nil pointer")
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	if !Swap(&a, &b) {
		t.Fatal("Swap should accept non-nil pointers")
	}
	if a != 2 || b != 1 {
		t.Fatalf("Swap result = (%d, %d), want (2, 1)", a, b)
	}
	if Swap(&a, nil) {
		t.Fatal("Swap should reject nil pointer")
	}
}

func TestLinkedNodes(t *testing.T) {
	head := Link(1, 2, 3)
	if got := Values(head); !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("Values() = %v, want [1 2 3]", got)
	}
}
