package OSExam

import "testing"

func TestFCFS_FCFS(t *testing.T) {
	fcfs := FCFS{}
	fcfs.InitFromFile("test.txt")

	fcfs.FCFS()
}