package OSExam

// 周转时间 = 作业完成时间 - 作业提交时间

type Process struct {
	finishTime float64
	submitTime float64
	startTime float64
	waitTime float64
	// 周转时间
	trTime float64
	wtrTime float64
	pid int
	// 这个是作业的实际运行时间
	runTime float64
	// 作业的优先级
	priority int
	// 标记是否访问过和是否到达
	reached bool
	visited bool
}

func (this *Process)GetColTime() {
	this.trTime = this.finishTime - this.submitTime
}

func (this *Process) GetColTimeWithWeight() {
	this.wtrTime = float64(this.trTime)/float64(this.runTime)
}