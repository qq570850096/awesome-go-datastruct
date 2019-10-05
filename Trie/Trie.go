package Trie
// 字典树(前缀树)

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
