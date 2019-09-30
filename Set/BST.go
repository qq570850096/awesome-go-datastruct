package Set

import "algo/BinarySearch"

type Set interface {
	Add (e int)
	Remove (e int)
	Contains (e int) bool
	GetSize () int
	IsEmpty () bool
}

type BST struct {
	bst *BinarySearch.Tree
}
// 添加元素
func (this *BST) Add(e int) {
	this.bst.AddE(e)
}
// 删除元素
func (this *BST) Remove(e int) {
	this.bst.Remove(e)
}
// 是否包含某个元素
func (this *BST) Contains(e int) bool {
	return this.bst.Contains(e)
}
// 返回size大小
func (this *BST) GetSize() int {
	return this.bst.Size()
}
// 返回集合是否为空
func (this *BST) IsEmpty() bool {
	return this.bst.IsEmpty()
}
