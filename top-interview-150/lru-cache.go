package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func ConstructorLRU(capacity int) LRUCache {
	l := LRUCache{
		size: 0, capacity: capacity,
		cache: make(map[int]*DLinkedNode),
		head:  new(DLinkedNode),
		tail:  new(DLinkedNode),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.cache[key]
	if !exist {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.cache[key]
	if !exist {
		insertNode := &DLinkedNode{key: key, value: value}
		this.addToHead(insertNode)
		this.cache[key] = insertNode
		this.size++

		if this.size > this.capacity {
			delNode := this.removeTail()
			delete(this.cache, delNode.key)
			this.size--
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	next := this.head.next
	next.prev = node
	node.next = next

	this.head.next = node
	node.prev = this.head
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	delete(this.cache, node.key)
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

//
//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		size: 0, capacity: capacity,
//		cache: make(map[int]*DLinkedNode),
//		head:  new(DLinkedNode), tail: new(DLinkedNode),
//	}
//	l.head.next = l.tail
//	l.tail.prev = l.head
//
//	return l
//}
//
//func (this *LRUCache) Get(key int) int {
//	node, exist := this.cache[key]
//	if !exist {
//		return -1
//	}
//	this.moveToHead(node)
//	return node.value
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	node, exist := this.cache[key]
//	if !exist {
//		newNode := &DLinkedNode{key: key, value: value}
//		this.addToHead(newNode)
//		this.size++
//		this.cache[key] = newNode
//
//		if this.size > this.capacity {
//			tail := this.removeTail()
//			delete(this.cache, tail.key)
//			this.size--
//		}
//		return
//	}
//
//	node.value = value
//	this.moveToHead(node)
//}
//
//func (this *LRUCache) addToHead(node *DLinkedNode) {
//	next := this.head.next
//	next.prev = node
//	node.next = next
//
//	this.head.next = node
//	node.prev = this.head
//}
//
//func (this *LRUCache) removeNode(node *DLinkedNode) {
//	node.prev.next = node.next
//	node.next.prev = node.prev
//}
//
//func (this *LRUCache) moveToHead(node *DLinkedNode) {
//	this.removeNode(node)
//	this.addToHead(node)
//}
//
//func (this *LRUCache) removeTail() *DLinkedNode {
//	node := this.tail.prev
//	this.removeNode(node)
//	return node
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
