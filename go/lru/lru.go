// LRU算法实现
package lru

type Node struct {
	key       string
	value     int
	pre, next *Node
}

// LRUCache ...
type LRUCache struct {
	count, capacity int
	head, tail      *Node
	cache           map[string]*Node
}

func New(capacity int) *LRUCache {
	res := new(LRUCache)
	res.count = 0
	res.capacity = capacity

	res.head = new(Node)
	res.head.pre = nil
	res.tail = new(Node)
	res.tail.next = nil
	res.head.next = res.tail
	res.tail.pre = res.head

	res.cache = map[string]*Node{}
	return res
}
func (lru *LRUCache) Get(key string) int {
	node, ok := lru.cache[key]
	if !ok {
		return -1
	}
	lru.moveNodeToHead(node)
	return node.value
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
func (lru *LRUCache) Set(key string, value int) {
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

func (lru *LRUCache) PrintCache() {
	node := lru.head.next
	for node != nil && node != lru.tail {
		println("key:", node.key, "\t value:", node.value)
		node = node.next
	}
	println("--------------------------------------")
}
