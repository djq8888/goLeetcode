package main

import "container/list"

type LRUCache struct {
	capacity int
	//链表中存储key
	cache *list.List
	//存储key-链表节点
	key2list map[int]*list.Element
	//存储链表节点-key
	list2key map[*list.Element]int
}


func Constructor(capacity int) LRUCache {
	lc := LRUCache{capacity:capacity, key2list: make(map[int]*list.Element), list2key:make(map[*list.Element]int)}
	lc.cache = list.New()
	return lc
}


func (this *LRUCache) Get(key int) int {
	if element, ok := this.key2list[key]; ok {
		this.cache.MoveToFront(this.key2list[key])
		return element.Value.(int)
	} else {
		return -1
	}
}


func (this *LRUCache) Put(key int, value int)  {
	if this.capacity == 0 {
		return
	}
	//如果是更新值
	if element, ok := this.key2list[key]; ok {
		element.Value = value
		this.cache.MoveToFront(element)
		return
	}
	//如果内存已满
	if this.cache.Len() == this.capacity {
		rmkey := this.list2key[this.cache.Back()]
		delete(this.key2list, rmkey)
		delete(this.list2key, this.cache.Back())
		this.cache.Remove(this.cache.Back())
	}
	this.cache.PushFront(value)
	this.key2list[key] = this.cache.Front()
	this.list2key[this.cache.Front()] = key
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
