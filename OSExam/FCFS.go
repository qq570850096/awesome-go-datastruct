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
		fmt.Printf("job %d arrived\n", i+1)
		fmt.Printf("arrival %f, service %f\n", this.pending[i].submitTime, this.pending[i].runTime)
		fmt.Printf("job is executing\n")
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("execution succeeded\n")
		if i == 0 {
			this.pending[this.ready[i]].finishTime = this.pending[this.ready[i]].runTime + this.pending[this.ready[i]].submitTime
			this.pending[this.ready[i]].waitTime = 0
		} else {

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
		fmt.Printf("wait %f\tturnaround %f\tweighted turnaround: %0.2f\n", this.pending[this.ready[i]].waitTime, this.pending[this.ready[i]].trTime, this.pending[this.ready[i]].wtrTime)
	}
}
