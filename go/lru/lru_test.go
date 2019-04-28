package lru

import "testing"

func TestLRUCache(t *testing.T) {
	lruCache := New(3)
	lruCache.Set("aaa", 1)
	lruCache.PrintCache()
	lruCache.Set("bbb", 2)
	lruCache.PrintCache()
	lruCache.Set("ccc", 3)
	lruCache.PrintCache()
	lruCache.Set("ddd", 4)
	lruCache.PrintCache()
	lruCache.Set("eee", 5)
	lruCache.PrintCache()
	println("这是访问了ddd后的结果")
	lruCache.Get("ddd")
	lruCache.PrintCache()
	lruCache.Set("fff", 6)
	lruCache.PrintCache()
	lruCache.Set("aaa", 7)
	lruCache.PrintCache()
}

func (lru *LRUCache) PrintCache() {
	node := lru.head.next
	for node != nil && node != lru.tail {
		println("key:", node.key.(string), "\t value:", node.value.(int))
		node = node.next
	}
	println("--------------------------------------")
}
