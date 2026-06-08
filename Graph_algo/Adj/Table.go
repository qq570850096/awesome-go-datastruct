package Adj

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Table struct {
	v   int
	e   int
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

func (table *Table) ReadFromFile(filename string) (err error) {
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
	table.v = v
	table.e = e

	table.adj = make([][]int, v)
	for i := 0; i < v; i++ {
		table.adj[i] = make([]int, 0)
	}

	for {
		if _, err = fmt.Fscanln(file, &v, &e); err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		if err = table.validateVertex(v); err != nil {
			return
		}
		if err = table.validateVertex(e); err != nil {
			return
		}
		if indexOfIntSlice(table.adj[v], e) != -1 {
			err = errors.New("Parallel edges are detected!")
			return
		}
		if v == e {
			err = errors.New("Self Loop is detected!")
			return
		}
		table.adj[v] = append(table.adj[v], e)
		table.adj[e] = append(table.adj[e], v)
	}
}

func (table *Table) validateVertex(check int) (err error) {
	if check >= table.v || check < 0 {
		err = errors.New("vertex is invalid")
	}
	return
}

func (table *Table) HasEdge(v int, e int) bool {
	if err := table.validateVertex(v); err != nil {
		panic(err)
	}
	if err := table.validateVertex(e); err != nil {
		panic(err)
	}
	return indexOfIntSlice(table.adj[v], e) != -1
}

func (table *Table) LinkedVertex(v int) []int {
	if err := table.validateVertex(v); err != nil {
		panic(err)
	}
	return table.adj[v]
}

func (table *Table) String() string {
	var (
		builder strings.Builder
		index   int
		value   []int

		finalValue int
	)
	fmt.Fprintf(&builder, "V = %d, E = %d", table.v, table.e)

	for index, value = range table.adj {
		fmt.Fprintf(&builder, "\n%d : ", index)
		for _, finalValue = range value {
			fmt.Fprintf(&builder, "%d ", finalValue)
		}
	}
	return builder.String()
}

func (table *Table) Degree(v int) (res int) {
	return len(table.LinkedVertex(v))
}
