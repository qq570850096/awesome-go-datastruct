package Adj

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 使用哈希存储表示无向图
type Hash struct {
	v   int
	e   int
	adj map[int][]int
}

func (hash *Hash) E() int {
	return hash.e
}

func (hash *Hash) SetE(e int) {
	hash.e = e
}

func (hash *Hash) V() int {
	return hash.v
}

func (hash *Hash) SetV(v int) {
	hash.v = v
}

// 从文件中读取并创建一个无向图
func (hash *Hash) ReadFromFile(filename string) (err error) {
	var (
		file *os.File
		v    int
		e    int
	)
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()
	// golang按照列读取每行数据的方法
	if _, err = fmt.Fscanln(file, &v, &e); err != nil {
		return
	}
	if v < 0 || e < 0 {
		return errors.New("图的顶点或边数不能小于0！")
	}
	hash.v = v
	hash.e = e

	// golang独特的建立二维数组的方法
	hash.adj = make(map[int][]int, v)
	for i := 0; i < v; i++ {
		hash.adj[i] = make([]int, 0)
	}
	for {
		if _, err = fmt.Fscanln(file, &v, &e); err != nil {
			return
		}
		if err = hash.ValidateVertex(v); err != nil {
			return
		}
		if err = hash.ValidateVertex(e); err != nil {
			return
		}
		if value, ok := hash.adj[v]; ok && Find(value, e) != -1 {
			err = errors.New("Parallel edges are detected!")
			return
		}
		if v == e {
			err = errors.New("Self Loop is detected!")
			return
		}
		hash.adj[v] = append(hash.adj[v], e)
		hash.adj[e] = append(hash.adj[e], v)
	}
	return
}

// 验证是否越界
func (hash *Hash) ValidateVertex(check int) (err error) {
	if check >= hash.v || check < 0 {
		err = errors.New("vertex is invalid")
	}
	return
}

// 检测一条边是否联通
func (hash *Hash) HasEdge(v int, e int) bool {
	if err := hash.ValidateVertex(v); err != nil {
		panic(err)
	}
	if err := hash.ValidateVertex(e); err != nil {
		panic(err)
	}
	if value, ok := hash.adj[v]; ok && Find(value, e) != -1 {
		return true
	}
	return false
}

// 根据一个顶点返回所有的联通边
func (hash *Hash) LinkedVertex(v int) (edgeArr []int) {
	edgeArr = make([]int, 0)
	if err := hash.ValidateVertex(v); err != nil {
		panic(err)
	}
	for _, value := range hash.adj[v] {
		edgeArr = append(edgeArr, value)
	}
	return
}

// 构造打印方法
func (hash *Hash) String() string {
	var (
		builder strings.Builder
		index   int
		value   []int
		//j int
		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d", hash.v, hash.e)
	// golang独特的遍历二维数组的方法
	for index, value = range hash.adj {
		fmt.Fprintf(&builder, "\n%d : ", index)
		for _, finalValue = range value {
			fmt.Fprintf(&builder, "%d ", finalValue)
		}
	}
	return builder.String()
}

// 计算一个节点的度数
func (hash *Hash) Degree(v int) (res int) {
	return len(hash.LinkedVertex(v))
}
