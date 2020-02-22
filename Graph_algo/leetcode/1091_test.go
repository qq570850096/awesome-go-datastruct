package leetcode

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	test := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	fmt.Println(shortestPathBinaryMatrix(test))
	cur := "0000"
	var curs []byte
	curs = []byte(cur)
	next := make([]string,0)
	for i,v := range curs  {
		curs[i] = ( v - '0' + 1 )%10 + '0'
		next = append(next,string(curs))

		curs[i] = v
		curs[i] = (v - '0' + 9) % 10 + '0'
		next = append(next,string(curs))
		curs[i] = v

	}
	fmt.Println(next)
}
