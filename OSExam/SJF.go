package OSExam

import (
	"fmt"
	"math"
	"os"
	"time"
)

type SJF struct {
	process []Process
}

func (this *SJF)InitFromFile(filename string)  {
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

func (this *SJF) FindNextSJF(finish float64) int {
	// p是已经到达且拥有最短运行时间的进程的下标
	// q是没有到达的进程中拥有最早到达时间的进程的下标
	var p,q int
	var minNeedTime,minReachTIme,minTime float64
	minNeedTime,minReachTIme,minTime = math.MaxInt32,math.MaxInt32,math.MaxInt32
	for i := 0; i < len(this.process); i++ {
		if !this.process[i].visited {
			// 第一种情况
			if this.process[i].submitTime <= finish && this.process[i].runTime <= minNeedTime {
				p = i
				minNeedTime = this.process[i].runTime
			} else if this.process[i].submitTime > finish && this.process[i].submitTime <= minReachTIme {
				if this.process[i].runTime < minTime {
					q = i
					minReachTIme = this.process[i].submitTime
					minTime = this.process[i].runTime
				}
			}
		}
	}
	if p != -1 {
		return p
	}
	return q
}

//短作业优先算法
func (this *SJF) SJF () {
	var i int
	//总的等待时间 //总的周转时间 //总的带权周转时间
	var wrTime,trTime,wtrTime,finish float64  
	finish = math.MaxInt32 //当前完成时间
	for i=0;i<len(this.process);i++ {
		finish= func(a,b float64) float64 {
			if a < b{
				return a
			}
			return b
		}(finish,this.process[i].submitTime)
	}
	fmt.Printf("短作业优先算法: \n")
	for i=0;i<len(this.process);i++ {
		index:=this.FindNextSJF(finish)
		fmt.Printf("第%d个作业",index+1)
		fmt.Printf("到达时间 %.2f,服务时间%.2f\n",
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