package basics

import "testing"

func TestZeroValuesAndConst(t *testing.T) {
	i, s, b := ZeroValues()
	if i != 0 || s != "" || b != false {
		t.Fatalf("unexpected zero values: %d %q %v", i, s, b)
	}
	if got := DoublePi(); got != 2*Pi {
		t.Fatalf("DoublePi = %v, want %v", got, 2*Pi)
	}
}

func TestControlFlow(t *testing.T) {
	if Max(1, 2) != 2 || Max(5, -1) != 5 {
		t.Fatalf("Max not working")
	}

	tests := map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	}
	for n, want := range tests {
		if got := FizzBuzz(n); got != want {
			t.Fatalf("FizzBuzz(%d) = %q, want %q", n, got, want)
		}
	}

	if TypeName(1) != "int" || TypeName("x") != "string" || TypeName(true) != "bool" {
		t.Fatalf("TypeName basic cases failed")
	}
	if TypeName(1.0) != "unknown" {
		t.Fatalf("TypeName(1.0) should be unknown")
	}
}

func TestFuncsAndClosure(t *testing.T) {
	if Sum() != 0 || Sum(1, 2, 3) != 6 {
		t.Fatalf("Sum unexpected result")
	}

	first, last := SplitName("John Doe")
	if first != "John" || last != "Doe" {
		t.Fatalf("SplitName failed: %q %q", first, last)
	}
	first, last = SplitName("Solo")
	if first != "Solo" || last != "" {
		t.Fatalf("SplitName single word failed: %q %q", first, last)
	}

	counter := NewCounter(0)
	if counter() != 1 || counter() != 2 {
		t.Fatalf("NewCounter not increasing as expected")
	}
}

