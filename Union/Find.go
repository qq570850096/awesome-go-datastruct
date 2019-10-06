package Union

import "errors"

type UF interface {
	Union(p,q int)
	IsConnect(p,q int) bool
	GetSize() int
}
// 只实现了QuickFind的并查集
type Find struct {
	id []int
}

func InitUnionFind(size int) *Find {
	unionFind :=  &Find{
		id:make([]int,size),
	}
	for i := range unionFind.id {
		unionFind.id[i] = i
	}
	return unionFind
}
func(this *Find) GetSize () int {
	return len(this.id)
}
// 查找元素p对应的元素编号
func(this *Find) Find (p int) (res int, err error) {
	if p < 0 || p >= len(this.id) {
		err = errors.New("p is out of range")
	} else {
		res = this.id[p]
	}
	return
}
// 合并p，q所属的集合
func(this *Find) Union(p,q int) {
	var (
		pid int
		qid int
		err error
	)
	pid,err = this.Find(p)
	if err != nil {
		panic(err)
	}
	qid,err = this.Find(q)
	if err != nil {
		panic(err)
	}
	if pid == qid {
		return
	}
	for i,v := range this.id {
		if v == pid {
			this.id[i] = qid
		}
	}
	return
}
// 查询元素p,q是否属与一个集合
func(this *Find) IsConnect (p,q int) bool {
	pid,_ := this.Find(p)
	qid,_ := this.Find(q)
	return pid == qid
}


type QuickFind struct {
	parent []int
}

func InitUnionQuickFind(size int) *QuickFind {
	unionFind :=  &QuickFind{
		parent:make([]int,size),
	}
	// 初始化时每个节点都是一颗树
	for i := range unionFind.parent {
		unionFind.parent[i] = i
	}
	return unionFind
}

func(this *QuickFind) GetSize () int {
	return len(this.parent)
}
// 查找元素p对应的元素编号，时间复杂度为O(h),h为树的高度
func(this *QuickFind) Find (p int) int {
	if p < 0 || p >= len(this.parent) {
		panic("p is out of range!")
	} else {
		for p != this.parent[p] {
			p = this.parent[p]
		}
	}
	return p
}
// 合并过程，时间复杂度为O(h),h为树的高度
func(this *QuickFind) Union(p,q int) {
	var (
		pid int
		qid int
	)
	pid = this.Find(p)
	qid = this.Find(q)

	if pid == qid {
		return
	}
	this.parent[pid] = qid
	return
}
// 查询元素p,q是否属与一个集合
func(this *QuickFind) IsConnect (p,q int) bool {
	return this.Find(p) == this.Find(q)
}