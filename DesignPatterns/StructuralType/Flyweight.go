package StructuralType

import "fmt"

type Flyweight interface {
	operate()
}

type ExamInfo struct {
	user string

	subject string
}

func (e *ExamInfo) operate() {
	fmt.Println(e.user, "starts exam, subject:", e.subject)
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

func (e *ExamInfo) String() string {
	return "ExamInfo{" +
		"user = " + e.user + "\n" +
		"subject = " + e.subject + "\n" + "}"
}

type ExamInfoFactory struct {
	pool map[string]*ExamInfo
}

func (e *ExamInfoFactory) GetExamInfo(subject string) (Ex *ExamInfo) {
	if e.pool == nil {
		e.pool = make(map[string]*ExamInfo)
	}
	if v, ok := e.pool[subject]; ok {
		Ex = v
		fmt.Println("fetch from pool", subject)
	} else {
		fmt.Println("create object and put into pool", subject)
		Ex = &ExamInfo{subject: subject}
		e.pool[subject] = Ex
	}
	return
}
