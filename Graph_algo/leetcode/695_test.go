package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T)  {

	test := [][]int{{1,1,0,1,1},{1,0,0,0,0},{0,0,0,0,1},{1,1,0,1,1}}
	fmt.Println(maxAreaOfIsland(test))
}
