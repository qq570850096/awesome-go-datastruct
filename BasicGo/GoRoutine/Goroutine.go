package GoRoutine

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func Thread() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Millisecond)
}

func ThreadWrong() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond*20)
}

func CounterWrong() int {
	counter := 0
	for i:=0; i<5000 ; i++ {
		// 这里我们开5000个协程但是会发现最好counter的结果总不是5000
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	return counter
}

func Counter() int {
	// 这里我们使用加锁的方式
	var mut sync.Mutex
	counter := 0
	for i:=0; i<5000 ; i++ {
		go func() {
			mut.Lock()
			// 一般加锁之后就直接下面defer还有Open一般就直接defer close
			defer mut.Unlock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	return counter
}

func WaitGroupExam() int {
	// 那么有没有一种方法，可以不用time.Sleep()这种笨方法呢
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i:=0; i<5000 ; i++ {
		// 每启动一个协程，就在waitgroup中添加一个
		wg.Add(1)
		go func() {
			// wg一般和锁或者通道一起用
			mut.Lock()
			defer mut.Unlock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

// CSP并发模式
func AsnyService() <-chan string {
	// 这里制造一个无缓冲区的channel，这种时候的通道是一个同步的操作
	// 也就是说只有收发双方都对接上了，才会进行传递
	// retCh := make(chan string)
	// 这里我们换成一个带缓冲区的channel，虽然长度为1，也可以让他变成异步执行
	retCh := make(chan string,1)
	go func() {
		ret := Service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func Service() string {
	time.Sleep(50*time.Millisecond)
	return "Done"
}

func otherTask()  {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond*50)
	fmt.Println("Task is Done")
}

func Producer(ch chan <-  int, group *sync.WaitGroup)  {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		defer close(ch)
		defer group.Done()
	}()
}

func Consumer(ch <-chan int,wg *sync.WaitGroup)  {
	go func() {
		defer wg.Done()
		for {
			// 如果ok是false，那么说明通道又关闭又没有剩余值了。
			if data,ok := <- ch ; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	}()
}

func Cancel() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		// 这里开启五个协程，当任务取消的时候break掉
		go func(i int,ctx context.Context) {

			defer wg.Done()
			for {
				if isCancelledWithctx(ctx) {
					break
				}
			}
			fmt.Println(i,"Cancelled")
		}(i,ctx)

	}
	cancel()
	wg.Wait()
}

func isCancelled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func isCancelledWithctx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// 使用buffered channel实现一个对象池
type ReusableObj struct {}

type ObjPool struct {
	// 用于缓冲可复用对象
	bufChan chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj,num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool)GetObj(timeout time.Duration) (*ReusableObj,error) {
	select {
	case ret := <-p.bufChan:
		return ret,nil
	case <-time.After(timeout): // 超时控制
		return nil,errors.New("time out!")
	}
}

func (p *ObjPool)ReleaseObj(obj *ReusableObj) error  {
	select {
	case p.bufChan <- obj :
		return nil
	default:
		return errors.New("Overflow!")
	}
}