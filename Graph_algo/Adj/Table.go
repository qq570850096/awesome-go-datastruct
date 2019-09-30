package Adj

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
// 使用邻接表表示无向图
type Table struct {
	v int
	e int
	adj [][]int
}

func (table *Table) E() int {
	return table.e
}

func (table *Table) SetE(e int) {
	table.e = e
}

func (table *Table) V() int {
	return table.v
}

func (table *Table) SetV(v int) {
	table.v = v
}

// 从文件中读取并创建一个无向图
func (table *Table) ReadFromFile (filename string)(err error){
	var (
		file *os.File
		v int
		e int
	)
	if file,err = os.Open(filename);err!=nil{
		return
	}
	defer file.Close()
	// golang按照列读取每行数据的方法
	if _, err = fmt.Fscanln(file, &v, &e);err!=nil{
		return
	}
	if v < 0 || e < 0 {
		return errors.New("图的顶点或边数不能小于0！")
	}
	table.v = v
	table.e = e

	// golang独特的建立二维数组的方法
	table.adj = make([][]int,v)
	for i:=0;i<v;i++{
		table.adj[i] = make([]int,0)
	}

	for {
		if _, err = fmt.Fscanln(file,&v,&e);err!=nil{
			return
		}
		if err = table.validateVertex(v);err!=nil {
			return
		}
		if err = table.validateVertex(e);err!=nil {
			return
		}
		if Find(table.adj[v],e) != -1 {
			err = errors.New("Parallel edges are detected!")
			return
		}
		if v == e{
			err = errors.New("Self Loop is detected!")
			return
		}
		table.adj[v] = append(table.adj[v], e)
		table.adj[e] = append(table.adj[e], v)
	}
	return
}
// 验证是否越界
func (table *Table) validateVertex(check int)(err error){
	if check>=table.v || check<0{
		err = errors.New("vertex is invalid")
	}
	return
}

// 检测一条边是否联通
func (table *Table) HasEdge(v int,e int) bool{
	if err := table.validateVertex(v);err!=nil{
		panic(err)
	}
	if err := table.validateVertex(e);err!=nil{
		panic(err)
	}
	return Find(table.adj[v],e)!=-1
}
// 根据一个顶点返回所有的联通边
func (table *Table) LinkedVertex(v int) []int{
	if err := table.validateVertex(v);err!=nil{
		panic(err)
	}
	return table.adj[v]
}
// 构造打印方法
func (table *Table) String() string {
	var (
		builder strings.Builder
		index int
		value []int
		//j int
		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d",table.v,table.e)
	// golang独特的遍历二维数组的方法
	for index,value = range table.adj {
		fmt.Fprintf(&builder,"\n%d : ",index)
		for _,finalValue = range value {
			fmt.Fprintf(&builder, "%d ",finalValue)
		}
	}
	return builder.String()
}
// 计算一个节点的度数
func (table *Table) Degree(v int)(res int){
	return len(table.LinkedVertex(v))
}

func Find(a []int,v int)(index int){
	for index,_ = range a {
		if v == a[index]{
			return
		}
	}
	return -1
}