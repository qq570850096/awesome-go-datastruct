package Interfaces

import "testing"

func TestCompare(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		b    interface{}
		want int
	}{
		{name: "int greater", a: 3, b: 2, want: 1},
		{name: "int equal", a: 3, b: 3, want: 0},
		{name: "string less", a: "a", b: "b", want: -1},
		{name: "float greater", a: 1.5, b: 1.2, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.a, tt.b); got != tt.want {
				t.Fatalf("Compare(%v, %v) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestComparePanics(t *testing.T) {
	tests := []struct {
		name string
		a    interface{}
		b    interface{}
	}{
		{name: "different types", a: 1, b: "1"},
		{name: "unsupported type", a: true, b: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic")
				}
			}()
			Compare(tt.a, tt.b)
		})
	}
}
