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
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	success := scanner.Scan()
	start := time.Now()
	for success {
		word := scanner.Text()
		word = strings.ToLower(word)
		for len(word) > 0 && (word[0] < 'a' || word[0] > 'z') {
			word = word[1:]
		}
		for len(word) > 0 && (word[len(word)-1] < 'a' || word[len(word)-1] > 'z') {
			word = word[:len(word)-1]
		}
		trie.Push(word)
		success = scanner.Scan()
	}
	if success == false {

		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	end := time.Now()

	fmt.Println("Pride and Prejudice contains", trie.Size(), "words")
	fmt.Println("counting the book took: ", end.Sub(start))
}
