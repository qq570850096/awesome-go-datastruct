package Sorts

import (
	"slices"
	"testing"
)

func TestGradeSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{name: "empty", in: []int{}, want: []int{}},
		{name: "single", in: []int{100}, want: []int{100}},
		{name: "duplicates", in: []int{4, 66, 66, 67, 55, 55, 66, 99, 100, 4, 67}, want: []int{4, 4, 55, 55, 66, 66, 66, 67, 67, 99, 100}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := append([]int(nil), tt.in...)
			GradeSort(got)
			if !slices.Equal(got, tt.want) {
				t.Fatalf("GradeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
