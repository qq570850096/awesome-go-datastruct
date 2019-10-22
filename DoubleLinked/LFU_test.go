package DoubleLinked

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestInitLFUCahe(t *testing.T) {
	var (
		file *os.File
		err error
	)
	if file, err = os.Open("test.txt"); err != nil {
		return
	}
	defer file.Close()
	LFU := InitLFUCahe(3)
	for i:=0;i<3;i++ {
		var key int
		fmt.Fscanf(file,"%d",&key)
		LFU.Put(key,rand.Intn(100))
	}
	t.Log(LFU)
	for {
		var key int
		if _,err = fmt.Fscanf(file,"%d",&key);err!=nil{
			break
		} else {
			t.Log(LFU.Get(key))
		}
		t.Log(LFU.String())
	}
	t.Log("程序正常退出,一共发生",LFU.count,"次缺页中断")
}
