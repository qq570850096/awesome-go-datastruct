## Trie字典树

字典树（Trie）是一种很特别的树状信息检索数据结构，如同其名，它的构成就像一本字典，可以让你快速的进行字符插入、字符串搜索等。

[Trie](http://en.wikipedia.org/wiki/Trie) 一词来自 re**trie**val，发音为 /tri:/ "tree"，也有人读为 /traɪ/ "try"。

字典树设计的核心思想是空间换时间，所以数据结构本身比较消耗空间。但它利用了字符串的**共同前缀（Common Prefix）**作为存储依据，以此来节省存储空间，并加速搜索时间。Trie 的字符串搜索时间复杂度为 **O(m)**，m 为最长的字符串的长度，其查询性能与集合中的字符串的数量无关。其在搜索字符串时表现出的高效，使得特别适用于构建文本搜索和词频统计等应用。

**字典树的性质**

1. 根节点（Root）不包含字符，除根节点外的每一个节点都仅包含一个字符；
2. 从根节点到某一节点路径上所经过的字符连接起来，即为该节点对应的字符串；
3. 任意节点的所有子节点所包含的字符都不相同；

如下图的 Trie 树中包含了字符串集合 ["Joe", "John", "Johnny", "Jane", "Jack"]。

![img](http://exia.gz01.bdysite.com/uploads/big/9858f33f68e7557278c34559ba8a5de9.jpg)

Trie 关键词查找过程：

1. 每次从根结点开始搜索；
2. 获取关键词的第一个字符，根据该字符选择对应的子节点，转到该子节点继续检索；
3. 在相应的子节点上，获取关键词的第二个字符，进一步选择对应的子节点进行检索；
4. 以此类推，进行迭代过程；
5. 在某个节点处，关键词的所有字母已被取出，则读取附在该节点上的信息，查找完成。

### 字典和字典树的查询能力比较

字典：如果有n个条目，使用树结构，查询的时间复杂度是O(logn)，估算一下如果有100万个条目（2^20) logn大约为20

Trie：查询每个条目的时间复杂度，和字典中一共有多少条目无关，他的时间复杂度为O(w)，w为查询单词的长度。

### Trie的优缺点

优点：

1. 插入和查询的效率很高，都为O(m)O(m)，其中 mm 是待插入/查询的字符串的长度。
   1. 关于查询，会有人说 hash 表时间复杂度是O(1)O(1)不是更快？但是，哈希搜索的效率通常取决于 hash 函数的好坏，若一个坏的 hash 函数导致很多的冲突，效率并不一定比Trie树高
2. Trie树中不同的关键字不会产生冲突。

3. Trie树只有在允许一个关键字关联多个值的情况下才有类似hash碰撞发生。

4. Trie树不用求 hash 值，对短字符串有更快的速度。通常，求hash值也是需要遍历字符串的。

5. Trie树可以对关键字按字典序排序。

缺点

1. 当 hash 函数很好时，Trie树的查找效率会低于哈希搜索。

2. 空间消耗比较大。

### Trie的数据结构

```go
// 考虑到不同语言不同情境，每个指针应该由若干指向下个节点的指针

type Node struct {
	isWord bool
	next map[byte]*Node
}

func InitNode(isWord bool) *Node {
	return &Node{
		isWord:isWord,
		next:make(map[byte]*Node),
	}
}
func InitNodeWithoutPram() *Node {
	return &Node{
		isWord:false,
		next:make(map[byte]*Node),
	}
}
type Trie struct {
	root *Node
	size int
}

func (t Trie) Size() int {
	return t.size
}

func InitTrie() *Trie {
	return &Trie{
		root:InitNodeWithoutPram(),
		size:0,
	}
}
```

### Trie中添加新元素

```go
// 向Trie中添加一个新的单词word
func (this *Trie)Push(word string)  {
	cur := this.root
	for _,v := range word {
		_,ok:= cur.next[byte(v)]
		if !ok {
			cur.next[byte(v)] = InitNodeWithoutPram()
		}
		cur = cur.next[byte(v)]
	}
	// 判断新来的单词是否之前添加过，没添加过才添加进去
	if cur.isWord == false {
		cur.isWord = true
		this.size++
	}
}
```

### Trie中查找一个单词是否存在

```go
func (this *Trie) Contains(word string) bool {
	cur := this.root
	for _,v := range word {
		if res,ok := cur.next[byte(v)];ok {
			cur = res
		} else {
			return false
		}
	}
	// 走完了也不一定就一定是true，比如panda和pan我们有panda，pan不一定是我们添加的单词
	return cur.isWord
}
```

### Trie中搜索前缀是否存在

```go
func (this *Trie) SearchPrefix (prefix string) bool {
	cur := this.root
	for _,v := range prefix {
		if res,ok := cur.next[byte(v)];ok {
			cur = res
		} else {
			return false
		}
	}
	return true
}
```

完成以上功能后，LeetCode[第208号问题](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)即可轻松通过了

### Trie和简单的模式匹配

```go
func (this *Trie) SearchPrefix (prefix string) bool {
	cur := this.root
	for _,v := range prefix {
		if res,ok := cur.next[byte(v)];ok {
			cur = res
		} else {
			return false
		}
	}
	return true
}

func (this *Trie) MatchSearch (word string) bool {
	return this.match(this.root,word,0)
}

func (this *Trie) match (node *Node,word string,index int) bool {
	if index == len(word) {
		return node.isWord
	}

	v := word[index]
	if byte(v) != '.' {
		if res,ok := node.next[byte(v)];ok {
			return this.match(res,word,index+1)
		} else {
			return false
		}
	} else {
		for k,_ := range node.next {
			if this.match(node.next[k],word,index+1) {
				return true
			}
		}
		return false
	}
}
```

这个功能完成之后LeetCode[第211号问题](https://leetcode-cn.com/problems/add-and-search-word-data-structure-design/)也可以通过了

### Trie和字符串映射

这道题是LeetCode[第677号问题](https://leetcode-cn.com/problems/map-sum-pairs/submissions/)

```go
type Node struct {
	val int
	next map[byte]*Node
}
func InitNode(val int) *Node {
	return &Node{
		val:val,
		next:make(map[byte]*Node),
	}
}
func InitNodeWithoutPram() *Node {
	return &Node{
		next:make(map[byte]*Node),
	}
}
type MapSum struct {
    root *Node
	size int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{
		root:InitNodeWithoutPram(),
		size:0,
	}
}


func (this *MapSum) Insert(word string, val int)  {
    cur := this.root
	for _,v := range word {
		_,ok:= cur.next[byte(v)]
		if !ok {
			cur.next[byte(v)] = InitNodeWithoutPram()
		}
		cur = cur.next[byte(v)]
	}
	// 判断新来的单词是否之前添加过，没添加过才添加进去
    cur.val = val
    this.size++
}


func (this *MapSum) Sum(prefix string) int {
    cur := this.root
	for _,v := range prefix {
		if res,ok := cur.next[byte(v)];ok {
			cur = res
		} else {
			return 0
		}
	}
    return this.sum(cur)
}
func (this *MapSum) sum (node *Node) int {
    res := node.val
    for k,_ := range node.next {
        res += this.sum(node.next[k])
    }
    return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
```

