package BehavioralType

import "fmt"

// 意思是:定义一个操作中的算法的框架，而将一些步骤延迟到子类中。 使得子类可以不改变-一个
//算法的结构即可重定义该算法的某些特定步骤。
//简单来说，就是为子类设计一个模板，以便在子类中可以复用这些方法。

// 以生活中.上班的过程为例，我们上班的通常流程是:起床洗漱->通勤(开车、坐公交、打车)
//->到达公司。从以上步骤可以看出，只有通勤部分是不一样的，其他都一样，因为开车可能会
//被限号，就只能打车或坐公交去公司了，下面我们用代码( 模板模式)来实现一下。


// 上班抽象模板接口
type AbstractWork interface {
	GotoWork(work AbstractWork)
	Getup()
	Commute()
	Arrive()
}

type AbsClass struct {}

func (a AbsClass) GotoWork(work AbstractWork) {
	a.Getup()
	work.Commute()
	a.Arrive()
}

func (a AbsClass) Getup() {
	fmt.Println("1. 起床")
}

func (a AbsClass) Commute() {}

func (a AbsClass) Arrive() {
	fmt.Println("3. 到达")
}

type DriveToWork struct {
	AbsClass
}


func (d *DriveToWork) Commute() {
	fmt.Println("2. 开车去公司")
}

func (d *DriveToWork) GotoWork(work AbstractWork){
	d.AbsClass.GotoWork(d)
}

type BusToWork struct {
	AbsClass
}

func (d *BusToWork) Commute() {
	fmt.Println("2. 坐公交去公司")
}

func (d *BusToWork) GotoWork(work AbstractWork) {
	d.AbsClass.GotoWork(d)
}

