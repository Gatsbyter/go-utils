// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的

package leecode

// 这是我的憨批算法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	now := l

	for l1 != nil && l2 != nil {
		tmp := new(ListNode)
		if l1.Val > l2.Val {
			tmp.Val = l2.Val
			l2 = l2.Next
		} else {
			tmp.Val = l1.Val
			l1 = l1.Next
		}

		now.Next = tmp
		now = now.Next
	}

	if l1 == nil {
		now.Next = l2
	} else {
		now.Next = l1
	}

	return l.Next
}

// 空间复杂度更低的办法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	cur := l

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}

	return l.Next
}
