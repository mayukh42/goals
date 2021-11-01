package graphs

import (
	"log"
	// u "github.com/mayukh42/goals/utils"
)

var SETS = map[int][]int{
	1:  {4, 7},
	6:  {9},
	2:  {3, 6},
	11: {12},
	8:  {11},
	5:  {8, 10},
}

func TestUFDS1() *UFDS {
	uf := NewUFDS("a")
	// makeset
	s1, i1 := uf.MakeSet(1)
	s4, i4 := uf.MakeSet(4)
	s7, i7 := uf.MakeSet(7)

	as := []int{i1, i4, i7}
	log.Printf("all singletons: %v.%v", uf, as)

	log.Printf("s4: %v, %v", s4, uf.NodeAt(i4))

	// union
	s1 = uf.Union(s1, s4)
	s1 = uf.Union(s7, s1)

	log.Printf("s1.size = %d", s1.Size)
	log.Printf("1-(4,7): %v", uf)

	log.Printf("consistent? %v", uf.Consistent())

	// find
	r7 := uf.Find(s7)
	log.Printf("root(7) = %v", r7)
	r1 := uf.Find(s1)
	log.Printf("root(1) = %v", r1)

	return uf
}

func TestUFDS2() *UFDS {
	uf := NewUFDS("b")
	// makeset
	s2, i2 := uf.MakeSet(2)
	s3, i3 := uf.MakeSet(3)
	s6, i6 := uf.MakeSet(6)
	s9, i9 := uf.MakeSet(9)

	as := []int{i2, i3, i6, i9}
	log.Printf("all singletons: %v.%v", uf, as)

	log.Printf("s6: %v, %v", s6, uf.NodeAt(i6))

	// union
	s6 = uf.Union(s6, s9)
	s2 = uf.Union(s2, s3)
	s2 = uf.Union(s2, s6)

	log.Printf("s2.size = %d", s2.Size)
	log.Printf("2-(3,6-(9)): %v", uf)

	log.Printf("consistent? %v", uf.Consistent())

	// find
	r3 := uf.Find(s3)
	log.Printf("root(3) = %v", r3)
	r6 := uf.Find(s6)
	log.Printf("root(6) = %v", r6)
	r9 := uf.Find(s9)
	log.Printf("root(9) = %v", r9)

	return uf
}
