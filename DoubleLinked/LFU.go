package DoubleLinked

import (
	"fmt"
	"math"
	"strings"
)

type LFUNode struct {
	freq int
	node *Node
}

func InitLFUNode(key, value interface{}) *LFUNode {
	return &LFUNode{
		freq: 0,
		node: InitNode(key, value),
	}
}

type LFUCache struct {
	capacity int
	find     map[interface{}]*LFUNode
	freq_map map[int]*List
	size     int
	count    int
}

func InitLFUCahe(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		find:     map[interface{}]*LFUNode{},
		freq_map: map[int]*List{},
	}
}

// 更新节点的频率
func (this *LFUCache) updateFreq(node *LFUNode) {
	freq := node.freq
	// 删除
	node.node = this.freq_map[freq].Remove(node.node)
	if this.freq_map[freq].size == 0 {
		delete(this.freq_map, freq)
	}

	freq++
	node.freq = freq
	if _, ok := this.freq_map[freq]; !ok {
		this.freq_map[freq] = InitList(10)
	}
	this.freq_map[freq].Append(node.node)
}
func findMinNum(fmp map[int]*List) int {
	min := math.MaxInt32
	for key, _ := range fmp {
		min = func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}(min, key)
	}
	return min
}
func (this *LFUCache) Get(key interface{}) interface{} {
	node, ok := this.find[key]
	if !ok {
		fmt.Println("发生了一次缺页中断")
		this.count++
		return -1
	}
	this.updateFreq(node)
	return node.node.value
}

func (this *LFUCache) Put(key, value interface{}) {
	if this.capacity == 0 {
		return
	}
	// 命中缓存
	if _, ok := this.find[key]; ok {
		node := this.find[key]
		node.node.value = value
		this.updateFreq(node)
	} else {
		if this.capacity == this.size {
			// 找到一个最小的频率
			min_freq := findMinNum(this.freq_map)
			list := this.freq_map[min_freq]
			evicted := list.Pop()
			if evicted != nil {
				delete(this.find, evicted.key)
				if list.size == 0 {
					delete(this.freq_map, min_freq)
				}
				this.size--
			}
		}
		node := InitLFUNode(key, value)
		node.freq = 1
		this.find[key] = node
		if _, ok := this.freq_map[node.freq]; !ok {
			this.freq_map[node.freq] = InitList(math.MaxInt32)
		}
		node.node = this.freq_map[node.freq].Append(node.node)
		this.size++
	}
}

func (this *LFUCache) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, "*******************")
	for k, v := range this.freq_map {
		fmt.Fprintf(&builder, "Freq = %d\n", k)
		fmt.Fprintln(&builder, v.String())
	}
	fmt.Fprintln(&builder, "*******************")
	return builder.String()
}
