package Union

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFind(t *testing.T)  {
	uni := InitUnionFind(1000000)
	size := uni.GetSize()
	m := 100000
	start := time.Now()
	for i:=0;i<m;i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uni.Union(a,b)
	}
	for i:=0;i<m;i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uni.IsConnect(a,b)
	}
	end := time.Now()
	fmt.Println("对于十万级数据量，没有QuickFind的并查集共用时：",end.Sub(start))
	quni := InitUnionQuickFind(1000000)
	start = time.Now()
	for i:=0;i<m;i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		quni.Union(a,b)
	}
	for i:=0;i<m;i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		quni.IsConnect(a,b)
	}
	end = time.Now()
	fmt.Println("对于十万级数据量，使用了QuickFind的并查集共用时：",end.Sub(start))
}
