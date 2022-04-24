package main

import "fmt"

type Trie struct {
	words []string
}

func Constructor() Trie {
	return Trie{
		words: make([]string, 0),
	}
}

func (this *Trie) Insert(word string) {
	this.words = append(this.words, word)
}

func (this *Trie) Search(word string) bool {
	for _, w := range this.words {
		if w == word {
			return true
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	for _, w := range this.words {
		if len(w) < len(prefix) {
			continue
		} else if w[0:len(prefix)] == prefix {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Running main")
	obj := Constructor()
	word := "Something"
	prefix := "Some"
	obj.Insert(word)
	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
	fmt.Printf("Got %v\n", param_2)
	fmt.Printf("Got %v\n", param_3)
}
