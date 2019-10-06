package Red_Black

import (
	"algo/BinarySearch"
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRBTree(t *testing.T)  {
	m := 10000000
	rbtree := &Tree{
		root:nil,
		size:0,
	}
	startTime := time.Now()
	for i:=0;i<m;i++ {
		rbtree.Push(i,rand.Intn(math.MaxInt32))
	}
	for i:=0;i<m;i++ {
		rbtree.Contains(i)
	}
	endTime := time.Now()
	fmt.Println("对于千万级数据，红黑树共用时：",endTime.Sub(startTime))

	BST := BinarySearch.Tree{}
	startTime = time.Now()
	for i:=0;i<m;i++ {
		BST.AddE(rand.Intn(math.MaxInt32))
	}
	for i:=0;i<m;i++ {
		BST.Contains(rand.Intn(math.MaxInt32))
	}
	endTime = time.Now()
	fmt.Println("对于千万级数据，二分搜索树共用时：",endTime.Sub(startTime))
}
