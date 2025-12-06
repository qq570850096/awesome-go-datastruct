package skiplists

import "math/rand"

const (
	maxLevel int     = 16 // 足够应对 2^16 个元素
	p        float32 = 0.25
)

// Node 表示跳表中的节点。
type Node struct {
	Score   float64
	Value   interface{}
	forward []*Node
}

func newElement(score float64, value interface{}, level int) *Node {
	return &Node{
		Score:   score,
		Value:   value,
		forward: make([]*Node, level),
	}
}

// SkipList 表示一份跳表。
// SkipList 的零值即可直接使用。
type SkipList struct {
	header *Node // 头节点是一个哨兵节点
	len    int   // 当前跳表长度（不含头节点）
	level  int   // 当前跳表高度（不含头节点）
}

// New 创建一份空的跳表。
func New() *SkipList {
	return &SkipList{
		header: &Node{forward: make([]*Node, maxLevel)},
	}
}

// 返回长度
func (s *SkipList) Size() int {
	return s.len
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}

// Front 返回跳表中的第一个节点，可能为 nil。
func (sl *SkipList) Front() *Node {
	return sl.header.forward[0]
}

// Next 返回当前节点之后的第一个节点。
func (e *Node) Next() *Node {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

// 在跳表中查找给定分值对应的节点。
// 如果存在返回 (*Node, true)，否则返回 (nil, false)。
func (sl *SkipList) Search(score float64) (element *Node, ok bool) {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// 将分值和数据插入跳表，返回对应的节点指针。
func (sl *SkipList) Insert(score float64, value interface{}) *Node {
	update := make([]*Node, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	// 分值已存在则直接替换数据并返回
	if x != nil && x.Score == score {
		x.Value = value
		return x
	}

	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.forward[i] = update[i].forward[i]
		update[i].forward[i] = e
	}
	sl.len++
	return e
}

// 删除并返回给定分值的节点；不存在则返回 nil。
func (sl *SkipList) Delete(score float64) *Node {
	update := make([]*Node, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]

	if x == nil || x.Score != score {
		return nil
	}
	for i := 0; i < len(x.forward); i++ {
		if update[i].forward[i] != x {
			break
		}
		update[i].forward[i] = x.forward[i]
	}
	for sl.level > 1 && sl.header.forward[sl.level-1] == nil {
		sl.level--
	}
	sl.len--
	return x
}
