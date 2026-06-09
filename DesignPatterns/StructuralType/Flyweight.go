package StructuralType

import "fmt"

// Flyweight is the shared-object interface. In this example, ExamInfo instances
// are shared by subject.
type Flyweight interface {
	operate()
}

// ExamInfo stores intrinsic state in subject and extrinsic state in user. A
// stricter flyweight would keep user outside the shared object entirely.
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

// ExamInfoFactory owns the flyweight pool and returns shared ExamInfo values by
// subject.
type ExamInfoFactory struct {
	pool map[string]*ExamInfo
}

// GetExamInfo reuses an existing subject object or creates and caches a new one.
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
