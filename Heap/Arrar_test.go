package Heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T)  {
	n := 1000000
	ArrHeap := MaxHeap{
		arr:make([]int,0),
	}
	start := time.Now()
	for i:=0; i < n; i++ {
		ArrHeap.Add(rand.Intn(math.MaxInt32))
	}
	end := time.Now()
	// 生成一个一百万空间的数组把堆装回去，如果我们的算法无误，那么数组将是有序的

	test_arr := make([]int,1000000)
	for i:=0; i<n; i++ {
		test_arr[i] = ArrHeap.RemoveMax()
	}
	for i:=1;i<n;i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("大顶堆运行成功！",end.Sub(start))

	HeapArr := &MaxHeap{}
	for i:=0; i<n; i++ {
		test_arr[i] = rand.Intn(math.MaxInt32)
	}
	start = time.Now()
	HeapArr.InitHeapWithArray(test_arr)
	end = time.Now()
	for i:=0; i<n; i++ {
		test_arr[i] = HeapArr.RemoveMax()
	}
	for i:=1;i<n;i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("Heapify 运行成功！",end.Sub(start))
}