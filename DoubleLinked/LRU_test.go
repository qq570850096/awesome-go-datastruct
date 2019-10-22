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
		err error
	)
	if file, err = os.Open("test.txt"); err != nil {
		return
	}
	defer file.Close()
	LRU := InitLRU(3)
	for i:=0;i<3;i++ {
		var key int
		fmt.Fscanf(file,"%d",&key)
		LRU.Put(key,rand.Intn(100))
	}
	t.Log(LRU)
	for {
		var key int
		if _,err = fmt.Fscanf(file,"%d",&key);err!=nil{
			break
		} else {
			LRU.Get(key)
		}
		t.Log(LRU.String())
	}
	t.Log("程序正常退出,一共发生",LRU.count,"次缺页中断")
}
