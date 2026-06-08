package DoubleLinked

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestLRUCache(t *testing.T) {
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open("test.txt"); err != nil {
		return
	}
	defer file.Close()
	LRU := InitLRU(3)
	for i := 0; i < 4; i++ {
		var key int
		fmt.Fscanf(file, "%d", &key)
		LRU.Put(key, rand.Intn(100))
	}
	t.Log(LRU)
	for {
		var key int
		if _, err = fmt.Fscanf(file, "%d", &key); err != nil {
			break
		} else {
			LRU.Get(key)
		}
		t.Log(LRU.String())
	}
	t.Log("program exited normally, page faults: ", LRU.count, "page faults")
}
