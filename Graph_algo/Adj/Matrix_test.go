package Adj

import (
	"fmt"
	"testing"
)

func TestMatrix_Degree(t *testing.T) {
	m := &Matrix{}
	m.ReadFromFile("../g.txt")
	fmt.Println(m)
}
