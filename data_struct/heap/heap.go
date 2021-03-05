package heap

// 堆是一个完全二叉树
// 堆的实际实现是数组
// 大根堆：堆中每一个节点的值都必须大于等于其子树中每个节点的值。
// 小根堆：堆中每一个节点的值都必须小于等于其子树中每个节点的值。

// 堆的操作
// 1: 堆化 -> 将一个普通数组变成堆 O(n)
// 2: 删除顶端元素 -> 然后shift-down
// 3: 添加元素 -> 然后shift-up
// 4: 堆排序 -> O(nlogn)
// 5: shift-up
// 6: shift-down

type Heap []int

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
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
	*h = append(*h, v.(int))
}
