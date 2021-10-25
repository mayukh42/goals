package utils

import (
	"fmt"
)

type Any interface{}
type List []Any

type Tuple struct {
	First  Any
	Second Any
}

func NewTuple(a, b Any) *Tuple {
	return &Tuple{
		First:  a,
		Second: b,
	}
}

func (t *Tuple) String() string {
	return fmt.Sprintf("(%v, %v)", t.First, t.Second)
}

type Comparator func(u1, u2 Any) bool
type ComparatorFn func(a, b Any) int
type BinaryInt func(a, b Any) int

func CompareInt(a, b Any) int {
	var (
		ai, bi int
		ok     bool
	)
	if ai, ok = a.(int); !ok {
		panic("unknown type for arg 0")
	}
	if bi, ok = b.(int); !ok {
		panic("unknown type for arg 1")
	}

	// log.Printf("a, b = %d, %d", ai, bi)

	if ai < bi {
		return -1
	} else if ai > bi {
		return 1
	}

	return 0
}
