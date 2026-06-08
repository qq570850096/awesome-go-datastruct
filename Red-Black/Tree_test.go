package Red_Black

import (
	"algo/BinarySearch"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRBTree(t *testing.B) {
	m := 10000000
	rbtree := &Tree{
		root: nil,
		size: 0,
	}
	startTime := time.Now()
	for i := 0; i < m; i++ {
		rbtree.Push(i, rand.Intn(math.MaxInt32))
	}
	for i := 0; i < m; i++ {
		rbtree.Contains(i)
	}
	endTime := time.Now()
	fmt.Println("red-black tree on 10M items took: ", endTime.Sub(startTime))

	BST := BinarySearch.Tree{}
	startTime = time.Now()
	for i := 0; i < m; i++ {
		BST.AddE(rand.Intn(math.MaxInt32))
	}
	for i := 0; i < m; i++ {
		BST.Contains(rand.Intn(math.MaxInt32))
	}
	endTime = time.Now()
	fmt.Println("BST on 10M items took: ", endTime.Sub(startTime))
}
