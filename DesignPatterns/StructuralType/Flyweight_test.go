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
		sub := "科目" + strconv.Itoa(i)
		exF.GetExamInfo(sub)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			// 注意在此嵌套循环中，我们的Examinfo是由工厂返回出来的
			// 很显然我们只创建了两次这个对象，而不是像嵌套理所应当的六次
			Examinfo := exF.GetExamInfo("科目"+strconv.Itoa(j))
			Examinfo.SetUser("考生"+strconv.Itoa(i))
			fmt.Println(Examinfo)
			Examinfo.operate()
		}
	}
}
