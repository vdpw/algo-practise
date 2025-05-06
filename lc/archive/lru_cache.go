package archive

import "runtime/debug"

/*
设计 LRU 缓存策略:
On Get(key):
	1. Check if the key exists in the hash map.
		• If it doesn’t exist: Return an indicator (e.g., -1 or null).
		• If it exists:
			• Move the node to the front (head) of the linked list to mark it as most recently used.
			• Return the node’s value.

On Put(key, value):
	1. Check if the key already exists.
		• If it exists:
			• Update the node’s value.
			• Move the node to the front of the linked list.
		• If it doesn’t exist:
			• If the cache is at capacity:
				• Remove the tail node (least recently used item).
				• Remove its entry from the hash map.
			• Create a new node with the key and value.
			• Add it to the front of the linked list.
			• Add it to the hash map.

*/
// https://leetcode.cn/problems/lru-cache
type LRUCache struct {
	cache    map[int]*DoubleLinkedNode
	capacity int

	dummyHead *DoubleLinkedNode
	dummyTail *DoubleLinkedNode
}

type DoubleLinkedNode struct {
	key   int
	value int
	prev  *DoubleLinkedNode
	next  *DoubleLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := &LRUCache{
		cache:     make(map[int]*DoubleLinkedNode),
		capacity:  capacity,
		dummyHead: &DoubleLinkedNode{}, // dummyHead -> head -> second -> third -> tail -> dummyTail
		dummyTail: &DoubleLinkedNode{}, // dummyTail <- tail <- third <- second <- head <- dummyHead
	}
	l.dummyHead.next = l.dummyTail
	l.dummyTail.prev = l.dummyHead
	return *l
}

func (this *LRUCache) removeNode(node *DoubleLinkedNode) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) addToHead(node *DoubleLinkedNode) {
	head := this.dummyHead.next

	this.dummyHead.next = node
	node.prev = this.dummyHead

	head.prev = node
	node.next = head
}

func (this *LRUCache) getNode(key int) *DoubleLinkedNode {
	if node, ok := this.cache[key]; ok {
		this.removeNode(node)
		this.addToHead(node)
		return node
	}
	return nil
}

func (this *LRUCache) Get(key int) int {
	if node := this.getNode(key); node != nil {
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node := this.getNode(key); node != nil {
		node.value = value
		return
	}
	if len(this.cache) >= this.capacity {
		tail := this.dummyTail.prev
		k := tail.key
		this.removeNode(tail)
		delete(this.cache, k)
	}
	newNode := &DoubleLinkedNode{key: key, value: value}
	this.cache[key] = newNode
	this.addToHead(newNode)
}
func init() { debug.SetGCPercent(-1) }
