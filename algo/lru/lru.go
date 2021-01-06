package lru

import "container/list"

type LruCache struct {
	Capacity int
	Cache    map[string]*list.Element
	list     *list.List
}

func NewLruCache(capacity int) *LruCache {
	return &LruCache{
		Capacity: capacity,
		Cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

type Entry struct {
	Key   string
	Value string
}

func (l *LruCache) Put(key string, value string) {
	if elem, ok := l.Cache[key]; ok {
		elem.Value.(*Entry).Value = value
		l.list.MoveToFront(elem)
	} else {
		elem := l.list.PushFront(&Entry{Key: key, Value: value})
		l.Cache[key] = elem
	}
}
