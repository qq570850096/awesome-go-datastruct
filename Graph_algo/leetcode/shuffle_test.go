package leetcode

import "testing"

func TestShuffle_Run(t *testing.T) {
	s := Shuffle{
		N: 1000000,

		n: 54,
		m: 2,
	}

	s.Run()
}
