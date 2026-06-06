package lowlevel

import "testing"

func TestInspectLayout(t *testing.T) {
	info := InspectLayout()
	if info.Size == 0 {
		t.Fatal("layout size should be positive")
	}
	if info.Align == 0 {
		t.Fatal("layout alignment should be positive")
	}
	if info.CountOffset <= 0 {
		t.Fatalf("CountOffset = %d, want after first field", info.CountOffset)
	}
	if info.NameOffset <= info.CountOffset {
		t.Fatalf("NameOffset = %d, want after CountOffset %d", info.NameOffset, info.CountOffset)
	}
}
