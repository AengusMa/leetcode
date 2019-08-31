// LRU算法实现
package lru

import (
	"sync"
)

type Node struct {
	key       interface{}
	value     interface{}
	pre, next *Node
}

// LRUCache ...
type LRUCache struct {
	count, capacity int
	head, tail      *Node
	cache           map[interface{}]*Node
	lock            sync.Mutex
}

func New(capacity int) *LRUCache {
	res := &LRUCache{
		count:    0,
		capacity: capacity,
		head:     new(Node),
		tail:     new(Node),
		cache:    make(map[interface{}]*Node),
	}
	res.head.pre = nil
	res.tail.next = nil
	res.head.next = res.tail
	res.tail.pre = res.head
	return res
}

// Get 获取
func (lru *LRUCache) Get(key string) interface{} {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	node, ok := lru.cache[key]
	if !ok {
		return nil
	}
	lru.moveNodeToHead(node)

	return node.value
}

// Get 设置
func (lru *LRUCache) Set(key string, value int) {
	lru.lock.Lock()
	defer lru.lock.Unlock()
	node, ok := lru.cache[key]
	if !ok {
		tmpNode := new(Node)
		tmpNode.key = key
		tmpNode.value = value
		lru.cache[key] = tmpNode

		lru.addNodeToHead(tmpNode)
		lru.increase()
	} else {
		node.value = value
		lru.moveNodeToHead(node)
	}

}

func (lru *LRUCache) containsKey(key string) bool {
	return false
}

func (lru *LRUCache) remove(key string) interface{} {
	return nil
}
func (lru *LRUCache) size(key string) int {
	return 0
}

func (lru *LRUCache) clear(key string) int {
	return 0
}

// 移动节点到头节点
func (lru *LRUCache) moveNodeToHead(node *Node) {
	if lru.head.next == node {
		return
	}
	lru.removeNode(node)
	lru.addNodeToHead(node)
}
func (lru *LRUCache) removeNode(node *Node) {
	pre := node.pre
	next := node.next

	pre.next = next
	next.pre = pre
}

// 添加节点到头节点
func (lru *LRUCache) addNodeToHead(node *Node) {
	node.pre = lru.head
	node.next = lru.head.next

	lru.head.next.pre = node
	lru.head.next = node
}

func (lru *LRUCache) increase() {
	lru.count++
	if lru.count > lru.capacity {
		tail := lru.popTail()
		delete(lru.cache, tail.key)
		lru.count--
	}
}
func (lru *LRUCache) popTail() *Node {
	res := lru.tail.pre
	lru.removeNode(res)
	return res
}
