package collections

import (
	"fmt"
	"log"

	u "github.com/mayukh42/goals/utils"
)

func EnqueueTest() {
	xs := u.List{1, 2, 3, 4, 5}
	q := CreateQueueFromArray(xs)
	fmt.Printf("queue: %v\n", q)
}

func DequeueTest() {
	xs := u.List{1, 2, 3, 4, 5}
	q := CreateQueueFromArray(xs)
	fmt.Printf("queue: %v, size: %d\n", q, q.Size())
	for !q.IsEmpty() {
		e, _ := q.Dequeue()
		fmt.Printf("%v\n", e)
	}
	fmt.Printf("queue: %v, size: %d\n", q, q.Size())
}

func PushTest() {
	xs := u.List{1, 2, 3, 4, 5}
	s := CreateStackFromArray(xs)
	fmt.Printf("stack: %v\n", s)
}

func PopTest() {
	xs := u.List{1, 2, 3, 4, 5}
	s := CreateStackFromArray(xs)
	fmt.Printf("stack: %v, size: %d\n", s, s.Size())
	for !s.IsEmpty() {
		e, _ := s.Pop()
		fmt.Printf("%v\n", e)
	}
	fmt.Printf("stack: %v, size: %d\n", s, s.Size())
}

func FSQTest() {
	xs := u.List{8, 0, 2, 7, 1, 4, 9}
	q := NewFSQueue(4)
	p := 2
	for i, x := range xs {
		// this already has q.Dequeue()
		e, _ := q.Enqueue(x)
		log.Printf("add %d: %v", x, q)
		if e != nil {
			log.Printf("remove %v due to max size", e)
		}

		if i == p || i == p+1 {
			e, _ := q.Dequeue()
			log.Printf("remove %v: %v", e, q)
		}
	}
}
