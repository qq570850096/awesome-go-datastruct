package DFS

import (
	"algo/Graph_algo/Adj"
	"os"
	"path/filepath"
	"testing"
)

func TestCCCountsConnectedComponents(t *testing.T) {
	graph := readHashGraph(t, "6 3\n0 1\n1 2\n3 4\n")

	cc := &CC{}
	cc.Init(graph)

	if got := cc.Cccount(); got != 3 {
		t.Fatalf("Cccount() = %d, want 3", got)
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
