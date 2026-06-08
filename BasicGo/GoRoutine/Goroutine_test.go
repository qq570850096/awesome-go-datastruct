package GoRoutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestThread(t *testing.T) {

	select {
	case ret := <-AsnyService():
		t.Log(ret)

	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}

func TestConsumer(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan int)

	wg.Add(1)
	go Producer(ch, &wg)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Consumer(ch, &wg)
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
		if v, err := pool.GetObj(time.Second); err != nil {
			t.Error(err)
		} else {
			fmt.Println(i)
			if err = pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
}
