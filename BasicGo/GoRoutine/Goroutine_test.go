package GoRoutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	//Thread()
	//ThreadWrong()
	//t.Log(CounterWrong())
	//t.Log(Counter())
	//t.Log(WaitGroupExam())
	//ret := AsnyService()
	//otherTask()
	//fmt.Println(<-ret)
	select {
	case ret := <-AsnyService():
		t.Log(ret)
	// 超时控制
	case <- time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

func TestConsumer(t *testing.T) {
	var wg sync.WaitGroup
	// 创造一个同步通道
	ch := make(chan int)
	// 创造一个异步通道
	// cha := make(chan int,3)
	wg.Add(1)
	go Producer(ch,&wg)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Consumer(ch,&wg)
	}
	t.Log(runtime.NumGoroutine())
	wg.Wait()
}

func TestCancel(t *testing.T) {
	Cancel()
}

func TestNewObjPool(t *testing.T) {
	pool := NewObjPool(10)

	for i := 0; i < 100; i++ {
		if v,err := pool.GetObj(time.Second);err!=nil{
			t.Error(err)
		} else {
			fmt.Println(i)
			if err = pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}