package data_struct

type Queue struct {
	q      []interface{}
	length int
}

func NewQueue() *Queue {
	queue := &Queue{
		q: make([]interface{}, 0, 2000),
	}
	return queue
}

func (q *Queue) EnQueue(val interface{}) {
	q.q = append(q.q, val)
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	val := q.q[0]
	q.q = q.q[1:]
	q.length--
	return val
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Empty() bool {
	return q.length == 0
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
