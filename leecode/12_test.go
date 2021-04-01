// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null
// 1、用map存已遍历过的节点，然后每次去map里比对 时空都是On
// 2、用快慢指针，挺牛逼的，脑筋急转弯
// 3、直接赌链表的地址从小到大分配，如下

package leecode

import (
	"testing"
	"unsafe"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test12(t *testing.T) {}

func detectCycle(head *ListNode) *ListNode {
	now := head

	for now != nil {
		if (uintptr(unsafe.Pointer(now.Next))) <= (uintptr(unsafe.Pointer(now))) {
			return now.Next
		}
		now = now.Next
	}

	return nil
}
