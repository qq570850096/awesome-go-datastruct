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
		// 调试变量：cc *DFS.CC
		cycleDetection *search.Cycle
		// 调试变量：Bip *search.BipartitionDetection
	)
	mar = &Adj.Hash{}
	// 初始化二分图检测器：Bip = &search.BipartitionDetection{}
	if err := mar.ReadFromFile("g2.txt"); err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println(mar)

	// 以下为单源路径调试代码：
	// 示例：single = new(search.SingleSource)
	// 示例：single.Init(mar,0)
	// 示例输出：fmt.Println(single.Pre())
	// 示例输出：fmt.Println("0 -> 6 : ",single.Path(6))
	// 构造路径示例：path = &search.Path{
	//	T:5, // 目标顶点
	//	SingleSource:single, // 源数据
	// } // 结构体示例结束
	// 示例输出：fmt.Println(path.Path())
	// 示例输出：fmt.Println(path.Visited())
	cycleDetection = new(search.Cycle)
	cycleDetection.Init(mar)
	fmt.Println(BFS.Traverse(mar, 0))
	// 输出是否存在环：fmt.Println(cycleDetection.HasCycle())
	// 初始化二分图检测：Bip.Init(mar)
	// 打印二分检测结果：fmt.Println(Bip.IsBippart())
}
