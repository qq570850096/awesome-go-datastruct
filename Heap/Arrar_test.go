package Heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T) {
	n := 1000000
	ArrHeap := MaxHeap{
		arr: make([]int, 0),
	}
	start := time.Now()
	for i := 0; i < n; i++ {
		ArrHeap.Add(rand.Intn(math.MaxInt32))
	}
	end := time.Now()

	test_arr := make([]int, 1000000)
	for i := 0; i < n; i++ {
		test_arr[i] = ArrHeap.RemoveMax()
	}
	for i := 1; i < n; i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("max heap completed", end.Sub(start))

	HeapArr := &MaxHeap{}
	for i := 0; i < n; i++ {
		test_arr[i] = rand.Intn(math.MaxInt32)
	}
	start = time.Now()
	HeapArr.InitHeapWithArray(test_arr)
	end = time.Now()
	for i := 0; i < n; i++ {
		test_arr[i] = HeapArr.RemoveMax()
	}
	for i := 1; i < n; i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("heapify completed", end.Sub(start))
}
