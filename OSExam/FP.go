package OSExam

import (
	"fmt"
	"math"
	"os"
	"time"
)

type FP struct {
	process []Process
	ready []int
}

func (this *FP) getReady ()  {
	this.ready = make([]int, len(this.process))
	for i:=0;i<len(this.process) ;i++  {
		this.ready[i] = i
	}
}

func (this *FP)Less (i,j int) bool {
	// 按照优先级进行递增排序
	return this.process[i].priority < this.process[j].priority
}

func (this *FP)Len() int {
	return len(this.process)
}

func (this *FP)Swap(i,j int)  {
	this.process[i],this.process[j] = this.process[j],this.process[i]
}

func (this *FP)InitFromFile(filename string)  {
	var (
		file *os.File
		id int
		sub float64
		run float64
		priority int
		err error
	)
	if file,err = os.Open(filename);err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		if _, err = fmt.Fscanln(file, &id, &sub,&run,&priority); err != nil {
			return
		}
		proce := Process{
			pid:id,
			submitTime:sub,
			runTime:run,
			priority:priority,
		}
		this.process = append(this.process,proce)
	}
	return
}

func (this *FP) FindNextHPF (time float64) int {
	// 优先值越低 表示优先权越高
	// p是已经到达且拥有最高优先权的进程的下标
	// q是没有到达的进程中拥有最早到达时间的进程的下标
	var i, p, q,minPrivilege1 int
	var minReachTime float64
	p = 0
	q = 0
	minPrivilege1 = math.MaxInt32
	minReachTime = math.MaxInt32
	for i = 0; i < len(this.process); i++  {
		if !this.process[i].visited {
			// 第一情况
			if this.process[i].submitTime<=time && this.process[i].priority <= minPrivilege1 {

				if this.process[i].priority==this.process[p].priority {   //如果优先权一致 则按最早抵达时间
					if this.process[i].submitTime<this.process[p].submitTime {
						p=i
					}
				} else {
					p = i
					minPrivilege1 = this.process[i].priority
				}
			} else if this.process[i].submitTime > time && this.process[i].submitTime <= minReachTime  {
				q = i
				minReachTime = this.process[i].submitTime
			}
		}
	}
	// p为-1时,代表在time时刻还没进程到达,此时选择下一个最早到达的进程q
	if p != -1 {
		return p
	}
	return q
}

func (this *FP) HPF ()  {
	var (
		i int
		finish,wrTime,trTime,wtrTime float64
	)
	finish = math.MaxInt32
	for i := 0; i < len(this.process); i++ {
		finish = func(a,b float64) float64 {
			if a < b {
				return a
			}
			return b
		}(finish,this.process[i].submitTime)
	}

	fmt.Printf("优先权高者优先服务(非抢占式): \n");
	for i=0;i<len(this.process);i++{
		index:=this.FindNextHPF(finish)
		fmt.Printf("第%d个作业",index+1)
		fmt.Printf("到达时间 --%.2f,服务时间--%.2f\n",
			this.process[index].submitTime,
			this.process[index].runTime)
		fmt.Printf("本作业正在运行...........\n")
		time.Sleep(100*time.Millisecond)
		fmt.Printf("运行完毕\n")
		if this.process[index].submitTime<=finish {
			this.process[index].waitTime=finish-this.process[index].submitTime
			this.process[index].startTime=finish
		} else{
			this.process[index].startTime=this.process[index].submitTime
			this.process[index].waitTime=0
		}
		this.process[index].finishTime=this.process[index].startTime+this.process[index].runTime
		this.process[index].trTime=this.process[index].finishTime-this.process[index].submitTime
		this.process[index].wtrTime=this.process[index].trTime/this.process[index].runTime
		this.process[index].visited=true
		wrTime+=this.process[index].waitTime
		trTime+=this.process[index].trTime
		wtrTime+=this.process[index].wtrTime
		finish=this.process[index].finishTime
		fmt.Printf("等待时间: %.2f 周转时间: %.2f 带权周转时间: %0.2f\n",this.process[index].waitTime,this.process[index].trTime,this.process[index].wtrTime)
	}
	fmt.Printf("--------所有作业调度完毕------\n")
	fmt.Printf("平均等待时间: %.2f 平均周转时间: %.2f 平均带权周转时间: %.2f",wrTime/float64(len(this.process)),trTime/float64(len(this.process)),wtrTime/float64(len(this.process)))
}
