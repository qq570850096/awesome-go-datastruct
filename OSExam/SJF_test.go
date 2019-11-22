package OSExam

import "testing"

func TestSJF(t *testing.T)  {
	sjf := &SJF{}
	sjf.InitFromFile("test1.txt")
	sjf.SJF()
}
