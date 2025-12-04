package OSExam

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type FCFS struct {
	pending []Process
	ready   []int
}

func (this *FCFS) Len() int {
	return len(this.pending)
}
func (this *FCFS) Less(i, j int) bool {
	return this.pending[i].submitTime < this.pending[j].submitTime
}
func (this *FCFS) Swap(i, j int) {
	this.pending[i] = this.pending[j]
}
func (this *FCFS) InitFromFile(filename string) {
	var (
		file *os.File
		id   int
		sub  float64
		run  float64
		err  error
	)
	if file, err = os.Open(filename); err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		if _, err = fmt.Fscanln(file, &id, &sub, &run); err != nil {
			return
		}
		proce := Process{
			pid:        id,
			submitTime: sub,
			runTime:    run,
		}
		this.pending = append(this.pending, proce)
	}
	return
}

func (this *FCFS) Push(process Process) {
	this.pending = append(this.pending, process)
}

func (this *FCFS) getReady() {
	this.ready = make([]int, len(this.pending))
	for i := 0; i < len(this.pending); i++ {
		this.ready[i] = i
	}
}

func (this *FCFS) FCFS() {
	sort.Sort(this)
	this.getReady()
	for i := 0; i < len(this.pending); i++ {
		fmt.Printf("到达第%d个作业\n", i+1)
		fmt.Printf("到达时间%f,服务时间%f\n", this.pending[i].submitTime, this.pending[i].runTime)
		fmt.Printf("本作业正在执行\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("执行成功！\n")
		if i == 0 {
			this.pending[this.ready[i]].finishTime = this.pending[this.ready[i]].runTime + this.pending[this.ready[i]].submitTime
			this.pending[this.ready[i]].waitTime = 0
		} else {
			//如果上一个作业的完成时间大于下一个作业的到达时间，则下一个作业的开始时间从上一个的完成时间开始
			if this.pending[this.ready[i-1]].finishTime > this.pending[this.ready[i]].submitTime {
				this.pending[this.ready[i]].finishTime = this.pending[this.ready[i-1]].finishTime + this.pending[this.ready[i]].runTime
				this.pending[this.ready[i]].waitTime = this.pending[this.ready[i-1]].finishTime - this.pending[this.ready[i]].submitTime
			} else {
				this.pending[this.ready[i]].finishTime = this.pending[this.ready[i]].runTime + this.pending[this.ready[i]].submitTime
				this.pending[this.ready[i]].waitTime = 0
			}
		}
		this.pending[this.ready[i]].GetColTime()
		this.pending[this.ready[i]].GetColTimeWithWeight()
		fmt.Printf("等待时间：%f\t周转时间：%f\t带权周转时间: %0.2f\n", this.pending[this.ready[i]].waitTime, this.pending[this.ready[i]].trTime, this.pending[this.ready[i]].wtrTime)
	}
}
