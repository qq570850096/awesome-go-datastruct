package OSExam

type Process struct {
	finishTime float64
	submitTime float64
	startTime  float64
	waitTime   float64

	trTime  float64
	wtrTime float64
	pid     int

	runTime float64

	priority int

	reached bool
	visited bool
}

func (this *Process) GetColTime() {
	this.trTime = this.finishTime - this.submitTime
}

func (this *Process) GetColTimeWithWeight() {
	this.wtrTime = float64(this.trTime) / float64(this.runTime)
}
