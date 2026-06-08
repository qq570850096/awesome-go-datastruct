package Set

import "algo/BinarySearch"

type Set interface {
	Add(e int)
	Remove(e int)
	Contains(e int) bool
	GetSize() int
	IsEmpty() bool
}

type BST struct {
	bst *BinarySearch.Tree
}

func InitBSTSet() *BST {
	return &BST{
		bst: &BinarySearch.Tree{},
	}
}

func (this *BST) Add(e int) {
	this.bst.AddE(e)
}

func (this *BST) Remove(e int) {
	this.bst.Remove(e)
}

func (this *BST) Contains(e int) bool {
	return this.bst.Contains(e)
}

func (this *BST) GetSize() int {
	return this.bst.Size()
}

func (this *BST) IsEmpty() bool {
	return this.bst.IsEmpty()
}
