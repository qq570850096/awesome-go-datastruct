package Trie

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestTrie(t *testing.T) {
	trie := InitTrie()
	file, err := os.Open("pride-and-prejudice.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// 缺省的分隔函数是bufio.ScanLines,我们这里使用ScanWords。
	// 也可以定制一个SplitFunc类型的分隔函数
	scanner.Split(bufio.ScanWords)
	// scan下一个token.
	success := scanner.Scan()
	start := time.Now()
	for success {
		word := scanner.Text()
		word = strings.ToLower(word)
		for len(word) > 0  && (word[0] < 'a' || word[0] > 'z') {
			word = word[1:]
		}
		for len(word) > 0  && (word[len(word)-1] < 'a' || word[len(word)-1] > 'z') {
			word = word[:len(word)-1]
		}
		trie.Push(word)
		success = scanner.Scan()
	}
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	end := time.Now()
	// fmt.Println(trie.SearchPrefix("pri"))
	fmt.Println("傲慢与偏见一共",trie.Size(),"个单词")
	fmt.Println("统计整本书共用时：",end.Sub(start))
}
