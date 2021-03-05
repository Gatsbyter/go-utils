package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	// list包里有两个主要类型List & Element
	// list包里的函数 就这一个 New一个list
	l := list.New() // 返回一个 *List

	for i := 1; i < 10; i++ {
		l.PushBack(i) // 往后面插入
	}

	e := l.PushFront(99) // 往前面插入 然后返回元素
	fmt.Println(e.Value)
	fmt.Println(e.Next().Value) // element就两个方法 next prev
	fmt.Println(e.Next().Prev().Value)

	l.InsertBefore(999, e)
	l.InsertAfter(888, e)

	l.Remove(e)

	// 遍历链表并打印它包含的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}
