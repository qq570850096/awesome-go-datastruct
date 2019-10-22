package DoubleLinked

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestFIFO(t *testing.T)  {

	var (
		file *os.File
		err error
	)
	if file, err = os.Open("test.txt"); err != nil {
		return
	}
	defer file.Close()
	Fifo := InitFIFO(3)
	for i:=0;i<3;i++ {
		var key int
		fmt.Fscanf(file,"%d",&key)
		Fifo.Put(key,rand.Intn(100))
	}
	t.Log(Fifo)
	for {
		var key int
		if _,err = fmt.Fscanf(file,"%d",&key);err!=nil{
			break
		} else {
			Fifo.Get(key)
		}
		t.Log(Fifo)
	}
	t.Log("程序正常退出,一共发生",Fifo.count,"次缺页中断")
}
