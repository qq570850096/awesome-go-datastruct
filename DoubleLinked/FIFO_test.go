package DoubleLinked

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestFIFO(t *testing.T) {
	all := 12.0
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open("test.txt"); err != nil {
		return
	}
	defer file.Close()
	Fifo := InitFIFO(4)
	for i := 0; i < 4; i++ {
		var key int
		fmt.Fscanf(file, "%d", &key)
		Fifo.Put(key, rand.Intn(100))
	}
	t.Log(Fifo)
	for {
		var key int
		if _, err = fmt.Fscanf(file, "%d", &key); err != nil {
			break
		} else {
			Fifo.Get(key)
		}
		t.Log(Fifo)
	}
	t.Log("program exited normally, page faults: ", Fifo.count, "page faults")
	t.Logf("page fault rate %.2f", float64(Fifo.count)/all)
}
