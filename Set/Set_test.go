package Set

import "testing"

func TestSetImplementations(t *testing.T) {
	tests := []struct {
		name string
		set  Set
	}{
		{name: "list", set: InitListSet()},
		{name: "bst", set: InitBSTSet()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.set.IsEmpty() {
				t.Fatal("new set should be empty")
			}

			tt.set.Add(3)
			tt.set.Add(1)
			tt.set.Add(3)

			if tt.set.GetSize() != 2 {
				t.Fatalf("duplicate add should keep size 2, got %d", tt.set.GetSize())
			}
			if !tt.set.Contains(1) || !tt.set.Contains(3) {
				t.Fatal("set should contain inserted values")
			}
			if tt.set.Contains(2) {
				t.Fatal("set should not contain missing value")
			}

			tt.set.Remove(3)
			if tt.set.Contains(3) {
				t.Fatal("removed value should be absent")
			}
			if tt.set.GetSize() != 1 {
				t.Fatalf("remove should reduce size to 1, got %d", tt.set.GetSize())
			}
		})
	}
}
