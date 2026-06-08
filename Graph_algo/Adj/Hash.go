package Adj

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

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

	if _, err = fmt.Fscanln(file, &v, &e); err != nil {
		return
	}
	if v < 0 || e < 0 {
		return errors.New("vertex or edge count cannot be negative")
	}
	hash.v = v
	hash.e = e

	hash.adj = make(map[int][]int, v)
	for i := 0; i < v; i++ {
		hash.adj[i] = make([]int, 0)
	}
	for {
		if _, err = fmt.Fscanln(file, &v, &e); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if err = hash.ValidateVertex(v); err != nil {
			return
		}
		if err = hash.ValidateVertex(e); err != nil {
			return
		}
		if value, ok := hash.adj[v]; ok && indexOfIntSlice(value, e) != -1 {
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
}

func (hash *Hash) ValidateVertex(check int) (err error) {
	if check >= hash.v || check < 0 {
		err = errors.New("vertex is invalid")
	}
	return
}

func (hash *Hash) HasEdge(v int, e int) bool {
	if err := hash.ValidateVertex(v); err != nil {
		panic(err)
	}
	if err := hash.ValidateVertex(e); err != nil {
		panic(err)
	}
	if value, ok := hash.adj[v]; ok && indexOfIntSlice(value, e) != -1 {
		return true
	}
	return false
}

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

func (hash *Hash) String() string {
	var (
		builder strings.Builder
		index   int
		value   []int

		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d", hash.v, hash.e)

	for index, value = range hash.adj {
		fmt.Fprintf(&builder, "\n%d : ", index)
		for _, finalValue = range value {
			fmt.Fprintf(&builder, "%d ", finalValue)
		}
	}
	return builder.String()
}

func (hash *Hash) Degree(v int) (res int) {
	return len(hash.LinkedVertex(v))
}
