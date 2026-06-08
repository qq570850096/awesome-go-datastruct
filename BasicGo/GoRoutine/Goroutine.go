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
	time.Sleep(time.Millisecond * 20)
}

func CounterWrong() int {
	counter := 0
	for i := 0; i < 5000; i++ {

		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second)
	return counter
}

func Counter() int {

	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			mut.Lock()

			defer mut.Unlock()
			counter++
		}()
	}
	time.Sleep(time.Second)
	return counter
}

func WaitGroupExam() int {

	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {

		wg.Add(1)
		go func() {

			mut.Lock()
			defer mut.Unlock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	return counter
}

func AsnyService() <-chan string {

	retCh := make(chan string, 1)
	go func() {
		ret := Service()
		fmt.Println("returned result")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func Service() string {
	time.Sleep(50 * time.Millisecond)
	return "Done"
}

func otherTask() {
	fmt.Println("Working on something else")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("Task is Done")
}

func Producer(ch chan<- int, group *sync.WaitGroup) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		defer close(ch)
		defer group.Done()
	}()
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {

			if data, ok := <-ch; ok {
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

		go func(i int, ctx context.Context) {

			defer wg.Done()
			for {
				if isCancelledWithctx(ctx) {
					break
				}
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)

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

type ReusableObj struct{}

type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, num)
	for i := 0; i < num; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out!")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("Overflow!")
	}
}
