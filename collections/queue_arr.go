package collections

import (
	"fmt"
	"strings"

	u "github.com/mayukh42/goals/utils"
)

// fixed size queue
type FSQueue struct {
	que   u.List
	size  int
	count int
	head  int
	tail  int
}

func NewFSQueue(size int) *FSQueue {
	return &FSQueue{
		que:   make(u.List, size),
		size:  size,
		count: 0,
		head:  0,
		tail:  0,
	}
}

func (fsq *FSQueue) Enqueue(elem u.Any) (u.Any, error) {
	var e u.Any = nil
	if fsq.count == fsq.size {
		// evict
		e, _ = fsq.Dequeue()
	}
	fsq.que[fsq.tail] = elem
	fsq.tail = (fsq.tail + 1) % fsq.size
	fsq.count++
	return e, nil
}

func (fsq *FSQueue) Dequeue() (u.Any, error) {
	e := fsq.que[fsq.head]
	fsq.head = (fsq.head + 1) % fsq.size
	fsq.count--
	return e, nil
}

func (fsq *FSQueue) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("fsq%d[%d] head: %v(%d) tail: %v(%d) \n",
		fsq.size, fsq.count,
		fsq.que[fsq.head], fsq.head,
		fsq.que[fsq.tail], fsq.tail,
	))
	// sb.WriteString(fmt.Sprintf("queue: %v \n", fsq.que))
	return sb.String()
}
