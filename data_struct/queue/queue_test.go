package queue

import "testing"

func TestQueue(t *testing.T) {
	queue := NewQueue()

	t.Logf("after new queue, queue: %v", queue)
	t.Logf("Is queue empty: %t", queue.Empty())
	t.Logf("length of queue: %d", queue.Length())
	t.Logf("Is queue has value %d: %t", 3, queue.HasValue(3))

	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	t.Logf("after enqueue 1 2 3 4, queue: %v", queue)
	t.Logf("Is queue empty: %t", queue.Empty())
	t.Logf("length of queue: %d", queue.Length())
	t.Logf("Is queue has value %d: %t", 3, queue.HasValue(3))
	t.Logf("first to list: %v", queue.ToList())

	t.Logf("first dequeue: %d", queue.Dequeue())
	t.Logf("second dequeue: %d", queue.Dequeue())

	t.Logf("after dequeue twice, queue: %v", queue)
	t.Logf("Is queue empty: %t", queue.Empty())
	t.Logf("length of queue: %d", queue.Length())
	t.Logf("Is queue has value %d: %t", 3, queue.HasValue(3))
	t.Logf("second to list: %v", queue.ToList())

	t.Logf("third dequeue: %d", queue.Dequeue())
	t.Logf("fourth dequeue: %d", queue.Dequeue())

	t.Logf("after dequeue fourth, queue: %v", queue)
	t.Logf("Is queue empty: %t", queue.Empty())
	t.Logf("length of queue: %d", queue.Length())
	t.Logf("Is queue has value %d: %t", 3, queue.HasValue(3))
	t.Logf("third to list: %v", queue.ToList())
}
