// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

package leecode

// 快慢指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	remove, tmp := head, head

	for i := 0; i < n; i++ {
		tmp = tmp.Next
	}

	if tmp == nil {
		return head.Next
	}

	for tmp.Next != nil {
		tmp = tmp.Next
		remove = remove.Next
	}

	remove.Next = remove.Next.Next
	return head
}
