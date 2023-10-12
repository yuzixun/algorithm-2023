package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {

	node := this
	for _, item := range word {
		key := item - 'a'
		if node.children[key] == nil {
			node.children[key] = &Trie{}
		}
		node = node.children[key]
	}

	node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, item := range prefix {
		node = node.children[item-'a']
		if node == nil {
			return nil
		}
	}

	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
