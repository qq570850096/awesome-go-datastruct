package OSExam

// 存储文件信息
type UFD struct {
	filename string
	attribute int //属性
	length int // 长度
	space [10]int // 为文件本身分配10个空间
	p [100]int	//一级索引，100个空间
	p2 [][]int  // 二级索引，100*100个空间
	next *UFD
}

// 存储目录信息
type DIR struct {
	above *DIR
	name string
	length int
	next *DIR
	FileHead *UFD //此目录下的文件指针
	DirHead *DIR // 此目录下目录链表指针
}

type Cuse struct {
	now *DIR
	FHead *UFD
	DHead *UFD

	username string
	password string
	length int // 用户空间大小
	status int // 是否获得空间
}

func(this *Cuse) SetStatus (int) {

}

func(this *Cuse) DeleteUser () {

}
func(this *Cuse) DisFile () {

}
func(this *Cuse) DisDir (d *DIR) {

}
