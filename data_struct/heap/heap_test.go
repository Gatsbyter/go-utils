package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &Heap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)

	fmt.Println(*h)
	fmt.Println(heap.Pop(h))
	heap.Push(h, 6)
	fmt.Println("new: ", *h)

	for len(*h) > 0 {
		fmt.Printf("%d \n ", heap.Pop(h))
	}
}

// 总结一下堆的方法
//func Init(h Interface)  // 对heap进行初始化，生成小根堆（或大根堆）
//func Push(h Interface, x interface{})  // 往堆里面插入内容 保证堆
//func Pop(h Interface) interface{}  // 从堆顶pop出内容 保证堆
//func Remove(h Interface, i int) interface{}  // 从指定位置删除数据，并返回删除的数据 保证堆
//func Fix(h Interface, i int)  // 从i位置数据发生改变后，对堆再平衡，优先级队列使用到了该方法
