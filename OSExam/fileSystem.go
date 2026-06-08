package OSExam

type UFD struct {
	filename  string
	attribute int
	length    int
	space     [10]int
	p         [100]int
	p2        [][]int
	next      *UFD
}

type DIR struct {
	above    *DIR
	name     string
	length   int
	next     *DIR
	FileHead *UFD
	DirHead  *DIR
}

type Cuse struct {
	now   *DIR
	FHead *UFD
	DHead *UFD

	username string
	password string
	length   int
	status   int
}

func (this *Cuse) SetStatus(int) {

}

func (this *Cuse) DeleteUser() {

}
func (this *Cuse) DisFile() {

}
func (this *Cuse) DisDir(d *DIR) {

}
