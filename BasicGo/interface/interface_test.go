package interfacedemo

import (
	"errors"
	"math"
	"testing"
)

func TestShapeInterface(t *testing.T) {
	shapes := []Shape{
		Rect{Width: 2, Height: 3},
		Circle{Radius: 1},
	}
	area := TotalArea(shapes)
	if area <= 0 {
		t.Fatalf("TotalArea should be positive, got %v", area)
	}
	wantRect := 2.0 * 3.0
	wantCircle := math.Pi * 1.0 * 1.0
	if diff := math.Abs(area - (wantRect+wantCircle)); diff > 1e-9 {
		t.Fatalf("unexpected area result: got %v, want %v", area, wantRect+wantCircle)
	}
}

func TestDescribe(t *testing.T) {
	if got := Describe(nil); got != "nil" {
		t.Fatalf("Describe(nil) = %q", got)
	}
	if got := Describe(3); got != "int:3" {
		t.Fatalf("Describe(3) = %q", got)
	}
	if got := Describe("go"); got != "string:go" {
		t.Fatalf("Describe(\"go\") = %q", got)
	}
	if got := Describe(true); got != "bool:true" {
		t.Fatalf("Describe(true) = %q", got)
	}
	if got := Describe(1.2); got != "unknown" {
		t.Fatalf("Describe(1.2) should be unknown, got %q", got)
	}
}

func TestErrorInterfaceAndWrapping(t *testing.T) {
	e := OpError{Op: "read", Code: 404, Msg: "not found"}
	if e.Error() == "" {
		t.Fatalf("OpError.Error should not be empty")
	}

	err := WrapAsTemporary("connect")
	if !errors.Is(err, ErrTemporary) {
		t.Fatalf("wrapped error should contain ErrTemporary, got %v", err)
	}
}

