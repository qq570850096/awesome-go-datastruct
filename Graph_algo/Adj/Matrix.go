package Adj

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Matrix struct {
	v   int
	e   int
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

func (matrix *Matrix) ReadFromFile(filename string) (err error) {
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
	matrix.v = v
	matrix.e = e

	matrix.adj = make([][]int, v)
	for i := 0; i < v; i++ {
		matrix.adj[i] = make([]int, v)
	}
	for {
		if _, err = fmt.Fscanln(file, &v, &e); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if err = matrix.validateVertex(v); err != nil {
			return
		}
		if err = matrix.validateVertex(e); err != nil {
			return
		}
		if matrix.adj[v][e] == 1 {
			err = errors.New("Parallel edges are detected!")
			return
		}
		if v == e {
			err = errors.New("Self Loop is detected!")
			return
		}
		matrix.adj[v][e] = 1
		matrix.adj[e][v] = 1
	}
}

func (matrix *Matrix) validateVertex(check int) (err error) {
	if check >= matrix.v || check < 0 {
		err = errors.New("vertex is invalid")
	}
	return
}

func (matrix *Matrix) HasEdge(v int, e int) bool {
	if err := matrix.validateVertex(v); err != nil {
		panic(err)
	}
	if err := matrix.validateVertex(e); err != nil {
		panic(err)
	}
	return matrix.adj[v][e] == 1
}

func (matrix *Matrix) LinkedVertex(v int) (edgearr []int) {
	edgearr = make([]int, 0)
	for i := 0; i < matrix.v; i++ {
		if matrix.adj[v][i] == 1 {
			edgearr = append(edgearr, i)
		}
	}
	return
}

func (matrix *Matrix) String() string {
	var (
		builder strings.Builder

		value []int

		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d", matrix.v, matrix.e)

	for _, value = range matrix.adj {
		fmt.Fprintf(&builder, "\n")
		for _, finalValue = range value {
			fmt.Fprintf(&builder, "%d ", finalValue)
		}
	}
	return builder.String()
}

func (matrix *Matrix) Degree(v int) (res int) {
	return len(matrix.LinkedVertex(v))
}
