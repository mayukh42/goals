package utils

import (
	"fmt"
	"strings"
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
		x, y int
		ok   bool
	)
	if x, ok = a.(int); !ok {
		panic("unknown type for arg 0")
	}
	if y, ok = b.(int); !ok {
		panic("unknown type for arg 1")
	}

	// log.Printf("a, b = %d, %d", ai, bi)

	if x < y {
		return -1
	} else if x > y {
		return 1
	}

	return 0
}

func CompareStr(a, b Any) int {
	var (
		x, y string
		ok   bool
	)
	if x, ok = a.(string); !ok {
		panic("unknown type for arg 0")
	}
	if y, ok = b.(string); !ok {
		panic("unknown type for arg 1")
	}

	return strings.Compare(x, y)
}

func AnyToString(a Any) string {
	return fmt.Sprintf("%v", a)
}
