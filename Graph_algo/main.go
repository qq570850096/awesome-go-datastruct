package main

import (
	"algo/Graph_algo/Adj"
	"algo/Graph_algo/BFS"
	"algo/Graph_algo/search"
	"fmt"
	"io"
)

func main() {
	var (
		mar *Adj.Hash

		cycleDetection *search.Cycle
	)
	mar = &Adj.Hash{}

	if err := mar.ReadFromFile("g2.txt"); err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(mar)

	cycleDetection = new(search.Cycle)
	cycleDetection.Init(mar)
	fmt.Println(BFS.Traverse(mar, 0))

}
