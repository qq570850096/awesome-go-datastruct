package OSExam

import "testing"

func TestFP_FP(t *testing.T) {
	fp := &FP{}
	fp.InitFromFile("test1.txt")
	fp.HPF()
}
