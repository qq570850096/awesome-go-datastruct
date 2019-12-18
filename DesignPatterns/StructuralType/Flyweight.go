package StructuralType

import "fmt"

// 享元模式是以共享的模式来支持大量细粒度对象的复用。听起来可能有点绕
// 其实java中的string就是一个享元模式
// String a = "abc";
// String b = "abc";
// System.out.pirntln(a==b) 会得到一个true
// 上面的例子中，a，b其实被创建的时候，都指向了常量池中某个字符串"abc"

//下面我们用一个最简单的考试报名的例子进行说明，假设我们有2个科目，有3位考生分别进
//行报考，我们一般会定义考试实体ExamInfo,如果不使用模式的话，可以想象，每次有考生参
//与一场科目考试的话，我们就会实例化一个ExamInfo,总共我们要实例化6个这样的实体，
//倘若使用享元模式，我们就只需要实例化2个这样的实体，然后通过内部状态的set方法进行
//不同对象的赋值操作，节省了不少的内存，很神奇吧?
type Flyweight interface {
	operate()
}


type ExamInfo struct {
	// 内部状态,用于在不同对象中共享
	user string
	
	// 外部状态，随环境改变而改变
	subject string
}

func (e *ExamInfo) operate() {
	fmt.Println(e.user,"开始考试，考试科目为",e.subject)
}

func (e *ExamInfo) User() string {
	return e.user
}

func (e *ExamInfo) SetUser(user string) {
	e.user = user
}

func (e *ExamInfo) Subject() string {
	return e.subject
}

func (e *ExamInfo) SetSubject(subject string) {
	e.subject = subject
}

func (e *ExamInfo) String () string {
	return "ExamInfo{" +
		"user = " + e.user + "\n" +
		"subject = " + e.subject + "\n" + "}" 
}

// 享元工厂
type ExamInfoFactory struct {
	pool map[string]*ExamInfo
}

func (e *ExamInfoFactory)GetExamInfo(subject string) (Ex *ExamInfo) {
	if v,ok := e.pool[subject]; ok {
		Ex = v
		fmt.Println("直接从池中获取",subject)
	} else {
		fmt.Println("建立对象，并且放入池中",subject)
		Ex = &ExamInfo{subject:subject}
		e.pool[subject] = Ex
	}
	return
}