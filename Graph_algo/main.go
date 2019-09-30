package main

import (
	"Graph_algo/Adj"
	"Graph_algo/BFS"
	"Graph_algo/search"
	"fmt"
	"io"
)

func main() {
	var (
		mar *Adj.Hash
		//cc *DFS.CC
		cycleDetection *search.Cycle
		//Bip *search.BipartitionDetection
	)
	mar = &Adj.Hash{}
	//Bip = &search.BipartitionDetection{}
	if err := mar.ReadFromFile("g2.txt"); err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(mar)

	//single = new(search.SingleSource)
	//single.Init(mar,0)
	//fmt.Println(single.Pre())
	//fmt.Println("0 -> 6 : ",single.Path(6))
	//path = &search.Path{
	//	T:5,
	//	SingleSource:single,
	//}
	//fmt.Println(path.Path())
	//fmt.Println(path.Visited())
	cycleDetection = new(search.Cycle)
	cycleDetection.Init(mar)
	fmt.Println(BFS.Traverse(mar, 0))
	//fmt.Println(cycleDetection.HasCycle())
	//Bip.Init(mar)
	//fmt.Println(Bip.IsBippart())
}
