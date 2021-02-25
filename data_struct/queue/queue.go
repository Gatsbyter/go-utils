package queue

type Queue struct {
	q []interface{}
}

func NewQueue() *Queue {
	queue := &Queue{
		q: make([]interface{}, 0, 2000),
	}
	return queue
}

func (q *Queue) EnQueue(val interface{}) {
	q.q = append(q.q, val)
}

func (q *Queue) Dequeue() interface{} {
	val := q.q[0]
	q.q = q.q[1:]
	return val
}

func (q *Queue) Length() int {
	return len(q.q)
}

func (q *Queue) Empty() bool {
	return len(q.q) == 0
}

func (q *Queue) HasValue(val interface{}) bool {
	for _, value := range q.q {
		if value == val {
			return true
		}
	}
	return false
}

func (q *Queue) ToList() []interface{} {
	return q.q
}
