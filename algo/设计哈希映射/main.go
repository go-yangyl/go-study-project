package main

import "container/list"

func main() {
	myHashMap := Constructor()
	myHashMap.Put(1, 1) // myHashMap 现在为 [[1,1]]
	myHashMap.Put(2, 2) // myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Get(1)    // 返回 1 ，myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Get(3)    // 返回 -1（未找到），myHashMap 现在为 [[1,1], [2,2]]
	myHashMap.Put(2, 1) // myHashMap 现在为 [[1,1], [2,1]]（更新已有的值）
	myHashMap.Get(2)    // 返回 1 ，myHashMap 现在为 [[1,1], [2,1]]
	myHashMap.Remove(2) // 删除键为 2 的数据，myHashMap 现在为 [[1,1]]
	myHashMap.Get(2)    // 返回 -1（未找到），myHashMap 现在为 [[1,1]]

}

const base = 32

type entry struct {
	key   int
	value int
}

type MyHashMap struct {
	data []list.List
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hash := hash(key)
	// 更新key
	for l := this.data[hash].Front(); l != nil; l = l.Next() {
		if et := l.Value.(*entry); et.key == key {
			et.value = value
			return
		}
	}

	// 插入key
	this.data[hash].PushBack(&entry{key: key, value: value})

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hash := hash(key)

	for l := this.data[hash].Front(); l != nil; l = l.Next() {
		if et := l.Value.(*entry); et.key == key {
			return et.value
		}
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hash := hash(key)

	for l := this.data[hash].Front(); l != nil; l = l.Next() {
		if et := l.Value.(*entry); et.key == key {
			this.data[hash].Remove(l)
		}
	}
}

func hash(key int) int {
	return key % base
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
