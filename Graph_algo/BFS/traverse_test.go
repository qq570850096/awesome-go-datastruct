package BFS

import (
	"algo/Graph_algo/Adj"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestTraverseVisitsAllComponents(t *testing.T) {
	graph := readHashGraph(t, "6 3\n0 1\n1 2\n3 4\n")

	got := Traverse(graph, 0)
	want := []int{0, 1, 2, 3, 4, 5}
	if !slices.Equal(got, want) {
		t.Fatalf("Traverse() = %v, want %v", got, want)
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
