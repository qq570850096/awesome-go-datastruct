package StructuralType

import (
	"fmt"
	"strconv"
	"testing"
)

func TestExamInfo_Subject(t *testing.T) {
	exF := &ExamInfoFactory{
		make(map[string]*ExamInfo),
	}
	for i := 0; i < 2; i++ {
		sub := "subject" + strconv.Itoa(i)
		exF.GetExamInfo(sub)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {

			Examinfo := exF.GetExamInfo("subject" + strconv.Itoa(j))
			Examinfo.SetUser("candidate" + strconv.Itoa(i))
			fmt.Println(Examinfo)
			Examinfo.operate()
		}
	}
}
