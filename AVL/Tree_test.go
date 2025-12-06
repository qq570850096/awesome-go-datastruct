package AVL

import (
	"reflect"
	"testing"
)

func TestAVLInsertBalanced(t *testing.T) {
	var tree Tree
	for i := 1; i <= 100; i++ {
		tree.Add(i)
	}
	if tree.Size() != 100 {
		t.Fatalf("size want 100, got %d", tree.Size())
	}
	if !tree.IsBST() {
		t.Fatalf("tree should keep BST order after inserts")
	}
	if !tree.IsBalanced() {
		t.Fatalf("tree should stay balanced after ordered inserts")
	}
	in := tree.InOrder()
	if len(in) != 100 || in[0] != 1 || in[len(in)-1] != 100 {
		sample := in
		if len(in) > 5 {
			sample = in[:5]
		}
		t.Fatalf("unexpected inorder result (sample): %#v", sample)
	}
}

func TestAVLRemoveKeepsBalance(t *testing.T) {
	var tree Tree
	vals := []int{30, 20, 40, 10, 25, 35, 50, 5, 15, 45, 60}
	for _, v := range vals {
		tree.Add(v)
	}
	for _, v := range []int{45, 10, 30, 60} {
		tree.Remove(v)
		if tree.Contains(v) {
			t.Fatalf("value %d should be removed", v)
		}
		if !tree.IsBST() {
			t.Fatalf("BST property broken after removing %d", v)
		}
		if !tree.IsBalanced() {
			t.Fatalf("tree unbalanced after removing %d", v)
		}
	}
}

func TestAVLDuplicatesAreIgnored(t *testing.T) {
	var tree Tree
	for _, v := range []int{5, 5, 5, 4, 4, 6} {
		tree.Add(v)
	}
	assertAVL(t, &tree, []int{4, 5, 6})
	if tree.Size() != 3 {
		t.Fatalf("size should count unique elements, got %d", tree.Size())
	}
}

func TestAVLRotations(t *testing.T) {
	// LL 型：左左失衡场景
	treeLL := &Tree{}
	for _, v := range []int{30, 20, 10} {
		treeLL.Add(v)
	}
	assertAVL(t, treeLL, []int{10, 20, 30})

	// RR 型：右右失衡场景
	treeRR := &Tree{}
	for _, v := range []int{10, 20, 30} {
		treeRR.Add(v)
	}
	assertAVL(t, treeRR, []int{10, 20, 30})

	// LR 型：左右失衡场景
	treeLR := &Tree{}
	for _, v := range []int{30, 10, 20} {
		treeLR.Add(v)
	}
	assertAVL(t, treeLR, []int{10, 20, 30})

	// RL 型：右左失衡场景
	treeRL := &Tree{}
	for _, v := range []int{10, 30, 20} {
		treeRL.Add(v)
	}
	assertAVL(t, treeRL, []int{10, 20, 30})
}

func TestAVLRemoveNonExisting(t *testing.T) {
	tree := &Tree{}
	for _, v := range []int{10, 20, 30} {
		tree.Add(v)
	}
	tree.Remove(999)
	assertAVL(t, tree, []int{10, 20, 30})
	if tree.Size() != 3 {
		t.Fatalf("size should remain unchanged when removing missing element, got %d", tree.Size())
	}
}

func TestAVLRemoveRootAndRebalance(t *testing.T) {
	tree := &Tree{}
	for _, v := range []int{50, 30, 70, 20, 40, 60, 80} {
		tree.Add(v)
	}
	tree.Remove(50)
	assertAVL(t, tree, []int{20, 30, 40, 60, 70, 80})
}

func assertAVL(t *testing.T, tree *Tree, want []int) {
	t.Helper()
	if !tree.IsBST() {
		t.Fatalf("BST property broken")
	}
	if !tree.IsBalanced() {
		t.Fatalf("tree is not balanced")
	}
	got := tree.InOrder()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("inorder mismatch, want %v, got %v", want, got)
	}
}
