package search

import (
	"algo/Graph_algo/Adj"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestSingleSourcePath(t *testing.T) {
	graph := readHashGraph(t, "5 3\n0 1\n1 2\n3 4\n")

	source := &SingleSource{}
	source.Init(graph, 0)

	if !source.IsConnectedTo(2) {
		t.Fatal("expected 0 to connect to 2")
	}
	if source.IsConnectedTo(4) {
		t.Fatal("expected 0 not to connect to 4")
	}
	if got, want := source.Path(2), []int{0, 1, 2}; !slices.Equal(got, want) {
		t.Fatalf("Path(2) = %v, want %v", got, want)
	}
	if got := source.Path(4); got != nil {
		t.Fatalf("Path(4) = %v, want nil", got)
	}
}

func TestCycleDetection(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    bool
	}{
		{name: "triangle", content: "3 3\n0 1\n1 2\n2 0\n", want: true},
		{name: "tree", content: "4 3\n0 1\n1 2\n1 3\n", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cycle := &Cycle{}
			cycle.Init(readHashGraph(t, tt.content))
			if got := cycle.HasCycle(); got != tt.want {
				t.Fatalf("HasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBipartitionDetection(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    bool
	}{
		{name: "square", content: "4 4\n0 1\n1 2\n2 3\n3 0\n", want: true},
		{name: "triangle", content: "3 3\n0 1\n1 2\n2 0\n", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bipartition := &BipartitionDetection{}
			bipartition.Init(readHashGraph(t, tt.content))
			if got := bipartition.IsBippart(); got != tt.want {
				t.Fatalf("IsBippart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readHashGraph(t *testing.T, content string) *Adj.Hash {
	t.Helper()
	filename := filepath.Join(t.TempDir(), "graph.txt")
	if err := os.WriteFile(filename, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	graph := &Adj.Hash{}
	if err := graph.ReadFromFile(filename); err != nil {
		t.Fatal(err)
	}
	return graph
}
