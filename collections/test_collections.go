package collections

import (
	"fmt"

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
