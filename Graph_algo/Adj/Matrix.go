package Adj

import (
	"errors"
	"fmt"
	"os"
	"strings"
)
// 使用邻接矩阵表示无向图
type Matrix struct {
	v int
	e int
	adj [][]int
}

func (matrix *Matrix) E() int {
	return matrix.e
}

func (matrix *Matrix) SetE(e int) {
	matrix.e = e
}

func (matrix *Matrix) V() int {
	return matrix.v
}

func (matrix *Matrix) SetV(v int) {
	matrix.v = v
}

// 从文件中读取并创建一个无向图
func (matrix *Matrix) ReadFromFile (filename string)(err error){
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
	matrix.v = v
	matrix.e = e

	// golang独特的建立二维数组的方法
	matrix.adj = make([][]int,v)
	for i:=0;i<v;i++{
		matrix.adj[i] = make([]int,v)
	}
	for {
		if _, err = fmt.Fscanln(file,&v,&e);err!=nil{
			return
		}
		if err = matrix.validateVertex(v);err!=nil {
			return
		}
		if err = matrix.validateVertex(e);err!=nil {
			return
		}
		if matrix.adj[v][e] == 1 {
			err = errors.New("Parallel edges are detected!")
			return
		}
		if v == e{
			err = errors.New("Self Loop is detected!")
			return
		}
		matrix.adj[v][e] = 1
		matrix.adj[e][v] = 1
	}
	return
}
// 验证是否越界
func (matrix *Matrix) validateVertex(check int)(err error){
	if check>=matrix.v || check<0{
		err = errors.New("vertex is invalid")
	}
	return
}

// 检测一条边是否联通
func (matrix *Matrix) HasEdge(v int,e int) bool{
	if err := matrix.validateVertex(v);err!=nil{
		panic(err)
	}
	if err := matrix.validateVertex(e);err!=nil{
		panic(err)
	}
	return matrix.adj[v][e]==1
}
// 根据一个顶点返回所有的联通边
func (matrix *Matrix) LinkedVertex(v int)(edgearr []int){
	edgearr = make([]int,0)
	for i:=0; i<matrix.v; i++ {
		if matrix.adj[v][i] == 1 {
			edgearr = append(edgearr,i)
		}
	}
	return
}
// 构造打印方法
func (matrix *Matrix) String() string {
	var (
		builder strings.Builder
		//index int
		value []int
		//j int
		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d",matrix.v,matrix.e)
	// golang独特的遍历二维数组的方法
	for _,value = range matrix.adj {
		fmt.Fprintf(&builder,"\n")
		for _,finalValue = range value {
			fmt.Fprintf(&builder, "%d ",finalValue)
		}
	}
	return builder.String()
}
// 计算一个节点的度数
func (matrix *Matrix) Degree(v int)(res int){
	return len(matrix.LinkedVertex(v))
}