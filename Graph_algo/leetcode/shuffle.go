package leetcode

import (
	"fmt"
	"math/rand"
	"time"
)

type Shuffle struct {
	N int

	n int

	m int
}

func (s Shuffle) Run() {

	rand.Seed(time.Now().Unix())
	freq := make([]int, s.n)
	arr := make([]int, s.n)
	for i := 0; i < s.N; i++ {
		reset(arr, s)
		shuffle(arr, s)
		for j := 0; j < s.n; j++ {
			freq[j] += arr[j]
		}
	}
	for i := 0; i < s.n; i++ {
		fmt.Println(i, ":", float64(freq[i])/float64(s.N))
	}
}

func reset(arr []int, s Shuffle) {

	for i := 0; i < s.m; i++ {
		arr[i] = 1
	}

	for i := s.m; i < s.n; i++ {
		arr[i] = 0
	}
}

func shuffle(arr []int, s Shuffle) {
	for i := 0; i < s.n; i++ {

		x := rand.Int()%(s.n-i) + i
		arr[i], arr[x] = arr[x], arr[i]
	}
}
