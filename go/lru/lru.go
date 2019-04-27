// LRU算法实现
package lru

type Node struct {
	key       string
	value     int
	pre, next Node
}

// LRUCache ...
type LRUCache struct {
	count, capacity int
	head, tail      Node
	cache           map[string]Node
}

func (lru *LRUCache) Get(key string) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	lru.moveNodeToHead(node)
	return node.value
}
func (lru *LRUCache) moveNodeToHead(node Node) {
	node.pre = lru.head
	node.next = lru.head.next

	lru.head.next.pre = node
	lru.head.next = node
}

func (lru *LRUCache) set(key string, value int) {
	node, ok := lru.cache[key]
	if !ok {

	} else {
		node.value = value
		lru.moveNodeToHead(node)
	}
}
