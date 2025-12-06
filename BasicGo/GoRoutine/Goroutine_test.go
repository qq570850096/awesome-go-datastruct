package GoRoutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestThread(t *testing.T) {
	// 调用 Thread()
	// 调用 ThreadWrong()
	// 打印 CounterWrong() 结果
	// 打印 Counter() 结果
	// t.Log(WaitGroupExam()) // 检查 WaitGroup 示例
	// ret := AsnyService() // 调用异步服务
	// otherTask() // 处理其他任务
	// fmt.Println(<-ret) // 读取异步结果
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
	// cha := make(chan int,3) // 带缓冲通道示例
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
