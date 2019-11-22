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

func InitLFUNode(key,value interface{}) *LFUNode {
	return &LFUNode{
		freq: 0,
		node: InitNode(key,value),
	}
}

type LFUCache struct {
	capacity int
	find map[interface{}]*LFUNode
	freq_map map[int]*List
	size int
	count int
}

func InitLFUCahe(capacity int) *LFUCache {
	return &LFUCache{
		capacity:capacity,
		find: map[interface{}]*LFUNode{},
		freq_map: map[int]*List{},
	}
}
// 更新节点的频率
func (this *LFUCache)updateFreq(node *LFUNode)  {
	freq := node.freq
	// 删除
	node.node = this.freq_map[freq].Remove(node.node)
	if this.freq_map[freq].size == 0 {
		delete(this.freq_map,freq)
	}

	freq++
	node.freq = freq
	if _,ok := this.freq_map[freq]; !ok {
		this.freq_map[freq] = InitList(10)
	}
	this.freq_map[freq].Append(node.node)
}
func findMinNum (fmp map[int]*List) int {
	min := math.MaxInt32
	for key,_ := range fmp {
		min = func(a,b int) int {
			if a > b {
				return b
			}
			return a
		}(min,key)
	}
	return min
}
func (this *LFUCache)Get(key interface{}) interface{} {
	if _,ok := this.find[key]; !ok {
		min_freq := findMinNum(this.freq_map)
		list := this.freq_map[min_freq]
		node := list.head
		fmt.Println("发生了一次缺页中断")
		// 先取到这个节点的地址
		newNode := this.find[node.key]
		// 从节点的映射中删掉这个节点
		delete(this.find,node.key)
		// 赋值新的key
		newNode.node.key = key
		this.find[key] = newNode
		// 在链表中真正的删除这个节点,并且删除后如果链表的长度为0,在频率映射表中吧这个链表删掉
		list.Remove(newNode.node)
		if list.size == 0{
			delete(this.freq_map,newNode.freq)
		}
		newNode.freq = 0
		if _,ok := this.freq_map[0]; !ok {
			this.freq_map[0] = InitList(10)
		}
		this.freq_map[0].Append(newNode.node)
		this.updateFreq(newNode)
		this.count++
		return -1
	}
	node := this.find[key]
	this.updateFreq(node)
	return node.node.value
}

func (this *LFUCache) Put (key,value interface{})  {
	if this.capacity == 0 {
		return
	}
	// 命中缓存
	if _,ok := this.find[key] ; ok {
		node := this.find[key]
		node.node.value = value
		this.updateFreq(node)
	} else {
		if this.capacity == this.size {
			// 找到一个最小的频率
			min_freq := findMinNum(this.freq_map)
			node := this.freq_map[min_freq].Pop()
			lfuNode := &LFUNode{
				node:node,
				freq:1,
			}
			this.find[key] = lfuNode
			delete(this.find,node.key)
			this.size--
		}
		node := InitLFUNode(key,value)
		node.freq = 1
		this.find[key] = node
		if _,ok := this.freq_map[node.freq]; !ok {
			this.freq_map[node.freq] = InitList(math.MaxInt32)
		}
		node.node = this.freq_map[node.freq].Append(node.node)
		this.size++
	}
}

func (this *LFUCache)String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder,"*******************")
	for k,v := range this.freq_map {
		fmt.Fprintf(&builder,"Freq = %d\n",k)
		fmt.Fprintln(&builder,v.String())
	}
	fmt.Fprintln(&builder,"*******************")
	return builder.String()
}