package link_list

import (
	"fmt"
	"testing"
)

func TestLinkList(t *testing.T) {

	head := New([]interface{}{1, 2, 3})

	head.InsertNext(4)
	head.InsertPrev(5)
	head.next.next.InsertNext(6).next.InsertPrev(7)
	Delete(head.next.next.next.next.next)

	t.Log(head.ToSlice())
}

type N struct {
	val  int
	next *N
}

func TestAAAAAA(t *testing.T) {
	head := BuildN(100)
	PrintN(head)
	head = RevertN(head)
	PrintN(head)
}

func RevertN(h *N) *N {
	head := h
	if head == nil || head.next == nil {
		return head
	}

	temp := head
	head = head.next
	next := head.next
	temp.next = nil

	for next != nil {
		head.next = temp
		temp = head
		head = next
		next = next.next
	}

	head.next = temp
	return head
}

func calLength(head *N) int {
	length := 0

	for {
		if head == nil {
			return length
		}

		length++
		head = head.next
	}
}

func BuildN(n int) *N {
	if n < 1 {
		return nil
	}

	head := &N{
		val:  1,
		next: nil,
	}
	temp := head

	for i := 2; i <= n; i++ {
		temp.next = &N{
			val:  i,
			next: nil,
		}
		temp = temp.next
	}

	return head
}

func PrintN(head *N) {
	for head != nil {
		fmt.Printf("%d -> ", head.val)
		head = head.next
	}
	fmt.Println()
}
