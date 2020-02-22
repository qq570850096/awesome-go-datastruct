package leetcode

import "testing"

func TestShuffle_Run(t *testing.T) {
	s := Shuffle{
		N: 1000000,
		// 计算一副牌中，大小王出现的概率应该为1/27也就是0.037
		n: 54,
		m: 2,
	}
	// 尝试运行时发现附和此概率
	s.Run()
}
