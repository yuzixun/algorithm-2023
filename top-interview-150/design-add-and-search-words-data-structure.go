package main

type WordDictionary struct {
	children map[byte]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children: make(map[byte]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	node := this

	for i := 0; i < len(word); i++ {
		if len(node.children) == 0 {
			node.children = make(map[byte]*WordDictionary)
		}
		if node.children[word[i]] == nil {
			node.children[word[i]] = &WordDictionary{children: make(map[byte]*WordDictionary)}
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

func (this *WordDictionary) searchWord(word string, index int, node *WordDictionary) bool {
	if node == nil {
		return false
	}
	if len(word) == index {
		return node.isEnd
	}
	if word[index] != '.' {
		return this.searchWord(word, index+1, node.children[word[index]])
	}
	for _, child := range node.children {
		if this.searchWord(word, index+1, child) {
			return true
		}
	}

	return false
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchWord(word, 0, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
