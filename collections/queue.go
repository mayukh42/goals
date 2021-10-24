package collections

import (
	"errors"

	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

type Queue struct {
	head *l.Node
	tail *l.Node
	size uint
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue) Enqueue(e u.Any) error {
	if e == nil {
		return errors.New("INVALID ELEMENT")
	}
	n := l.NewNode(e)
	if q.head == nil {
		q.head = n
	}
	q.tail = l.Link(q.tail, n)
	q.size++
	return nil
}

func (q *Queue) Dequeue() (u.Any, error) {
	if !q.IsEmpty() {
		n := q.head
		q.head = n.Next
		q.size--
		return n.Value, nil
	}
	return nil, errors.New("EMPTY QUEUE")
}

func (q *Queue) Size() uint {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) String() string {
	if q.head != nil {
		return q.head.String()
	}
	return "[]"
}

func CreateQueueFromArray(xs u.List) *Queue {
	q := NewQueue()
	for _, x := range xs {
		q.Enqueue(x)
	}
	return q
}
