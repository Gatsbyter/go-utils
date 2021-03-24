package leecode

import (
	"container/list"
	"testing"
)

func Test11(t *testing.T) {

}

type Item struct {
	value int
	ele   *list.Element
}

type LRUCache struct {
	l   *list.List
	m   map[int]Item
	cap int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		l:   list.New(),
		m:   make(map[int]Item),
		cap: capacity,
	}

	return l
}

func (l *LRUCache) Get(key int) int {
	val, ok := l.m[key]
	if !ok {
		return -1
	}

	l.l.Remove(val.ele)
	l.l.PushBack(key)
	l.m[key] = Item{
		value: l.m[key].value,
		ele:   l.l.Back(),
	}

	return val.value
}

// k1 -> k3 -> k2

func (l *LRUCache) Put(key int, value int) {
	val, ok := l.m[key]
	if ok {
		l.l.Remove(val.ele)
		l.l.PushBack(key)
		l.m[key] = Item{
			value: value,
			ele:   l.l.Back(),
		}
		return
	}

	if len(l.m) == l.cap {
		rm := l.l.Front()
		l.l.Remove(rm)
		delete(l.m, rm.Value.(int))
	}

	l.l.PushBack(key)
	l.m[key] = Item{
		value: value,
		ele:   l.l.Back(),
	}
}
