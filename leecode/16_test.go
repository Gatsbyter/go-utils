// 合并n的升序列表

package leecode

import (
	"container/heap"
)

// 优先级队列 很牛逼
func mergeKLists(lists []*ListNode) *ListNode {
	l := new(ListNode)
	cur := l

	h := &Heap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	for h.Len() > 0 {
		min := heap.Pop(h).(*ListNode)
		if min == nil {
			break
		}

		cur.Next = &ListNode{
			Val: min.Val,
		}
		cur = cur.Next
		if min.Next != nil {
			heap.Push(h, min.Next)
		}
	}

	return l.Next
}

type Heap []*ListNode

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}
