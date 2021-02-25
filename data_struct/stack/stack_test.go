package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()

	t.Log(stack.ToSlice())

	stack.Push(1)

	t.Log(stack.ToSlice())
	t.Log(stack.Peak())
	t.Log(stack.Pop())
	t.Log(stack.ToSlice())
	t.Log(stack.Pop())
	t.Log(stack.ToSlice())

	stack.Push(2)
	t.Log(stack.ToSlice())
	stack.Push(3)
	t.Log(stack.ToSlice())
	stack.Push(4)
	t.Log(stack.ToSlice())
	t.Log(stack.Pop())
	t.Log(stack.ToSlice())
}
