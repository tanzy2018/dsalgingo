package queue

import "errors"

type Queuer interface {
	Enqueue(interface{}) error
	Dequeue() (interface{}, error)
	Clear()
	Cap() int
	Len() int
}

// 非线性安全的队列
type Queue struct {
	queue []interface{}
	cap   int
	len   int
}

var (
	// 队列元素达到容量上限
	ErrQueueIsFull = errors.New("queue is up to it's capcity")
	// 队列中元素为空
	ErrQueueIsEmpty = errors.New("queue is empty")
)

func NewQueue(cap int) *Queue {
	if cap <= 0 {
		cap = 64
	}
	return &Queue{
		len:   0,
		cap:   cap,
		queue: make([]interface{}, 0, cap),
	}
}

func (q *Queue) Enqueue(v interface{}) error {
	if q.len == q.cap {
		return ErrQueueIsFull
	}
	q.queue = append(q.queue, v)
	q.len++
	return nil
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.len == 0 {
		return nil, ErrQueueIsEmpty
	}
	q.len--
	v := q.queue[q.len]
	q.queue = q.queue[:q.len]
	return v, nil
}

func (q *Queue) Clear() {
	if q.len > 0 {
		q.len = 0
		q.queue = make([]interface{}, 0, q.cap)
	}
}

func (q *Queue) Cap() int {
	return q.cap
}

func (q *Queue) Len() int {
	return q.len
}
