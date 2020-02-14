package Adj

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	m := &Table{}
	m.ReadFromFile("../g.txt")
	fmt.Println(m)
}
