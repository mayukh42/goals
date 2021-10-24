package collections

import (
	"errors"

	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

type Stack struct {
	head *l.Node
	size uint
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
		size: 0,
	}
}

func (s *Stack) Push(e u.Any) error {
	if e == nil {
		return errors.New("INVALID ELEMENT")
	}
	n := l.NewNode(e)
	l.Link(n, s.head)
	s.head = n
	s.size++
	return nil
}

func (s *Stack) Pop() (u.Any, error) {
	if !s.IsEmpty() {
		n := s.head
		s.head = n.Next
		s.size--
		return n.Value, nil
	}
	return nil, errors.New("EMPTY STACK")
}

func (s *Stack) Peek() (u.Any, error) {
	if !s.IsEmpty() {
		n := s.head
		return n.Value, nil
	}
	return nil, errors.New("EMPTY STACK")
}

func (s *Stack) Size() uint {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) String() string {
	if s.head != nil {
		return s.head.String()
	}
	return "[]"
}

func CreateStackFromArray(xs u.List) *Stack {
	s := NewStack()
	for _, x := range xs {
		s.Push(x)
	}
	return s
}
