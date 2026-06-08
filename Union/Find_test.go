package Union

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	uni := InitUnionFind(1000000)
	size := uni.GetSize()
	m := 100000
	start := time.Now()
	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uni.Union(a, b)
	}
	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uni.IsConnect(a, b)
	}
	end := time.Now()
	fmt.Println("union-find without QuickFind on 100k items took: ", end.Sub(start))
	quni := InitUnionQuickFind(1000000)
	start = time.Now()
	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		quni.Union(a, b)
	}
	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		quni.IsConnect(a, b)
	}
	end = time.Now()
	fmt.Println("union-find with QuickFind on 100k items took: ", end.Sub(start))
}
