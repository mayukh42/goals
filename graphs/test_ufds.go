package graphs

import (
	"log"
	// u "github.com/mayukh42/goals/utils"
)

// TODO: cleanup the tests or move the ref fns in another file
var SETS = map[int][]int{
	1:  {4, 7},
	6:  {9},
	2:  {3, 6},
	10: {12},
	8:  {11},
	5:  {8, 10},
}

func TestUFDS1() *UFDSRef {
	ufRef := NewUFDSRef("a-1")
	uf := ufRef.Set
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

	return ufRef
}

func TestUFDS2() *UFDSRef {
	ufRef := NewUFDSRef("b-2")
	uf := ufRef.Set
	// makeset
	s2, i2 := uf.MakeSet(2)
	s3, i3 := uf.MakeSet(3)
	s6, i6 := uf.MakeSet(6)
	s9, i9 := uf.MakeSet(9)

	bs := []int{i2, i3, i6, i9}
	log.Printf("all singletons: %v.%v", uf, bs)

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

	return ufRef
}

func TestUFDS3() *UFDSRef {
	ufRef := NewUFDSRef("c-5")
	uf := ufRef.Set
	// makeset
	s5, i5 := uf.MakeSet(5)
	s8, i8 := uf.MakeSet(8)
	s10, i10 := uf.MakeSet(10)
	s11, i11 := uf.MakeSet(11)
	s12, i12 := uf.MakeSet(12)

	cs := []int{i5, i8, i10, i11, i12}
	log.Printf("all singletons: %v.%v", uf, cs)

	log.Printf("s11: %v, %v", s11, uf.NodeAt(i11))

	// union
	s8 = uf.Union(s8, s11)
	s5 = uf.Union(s5, s10)
	s5 = uf.Union(s5, s8)
	s5 = uf.Union(s5, s12)

	log.Printf("s5.size = %d", s5.Size)
	log.Printf("5-(8-(11),10-(12)): %v", uf)

	log.Printf("consistent? %v", uf.Consistent())

	// find
	r8 := uf.Find(s8)
	log.Printf("root(8) = %v", r8)
	r11 := uf.Find(s11)
	log.Printf("root(11) = %v", r11)
	r12 := uf.Find(s12)
	log.Printf("root(12) = %v", r12)

	// after path compression
	log.Printf("5-(8-(11),10-(12)): %v", uf)

	// ufdsref
	// log.Printf("ufdsref: %v", ufRef)

	return ufRef
}

func TestUFDSRef1() *UFDSRef {
	ref := NewUFDSRef("b-2")
	// makeset
	ref.MakeSet(2)
	ref.MakeSet(3)
	ref.MakeSet(6)
	ref.MakeSet(9)

	log.Printf("all singletons: %v", ref)

	uf := ref.Set

	s6, i6 := ref.GetSingleton(6)
	log.Printf("s6: %v, %v", s6, uf.NodeAt(i6))

	s2, _ := ref.GetSingleton(2)
	s3, _ := ref.GetSingleton(3)
	s9, _ := ref.GetSingleton(9)

	// union
	s6 = uf.Union(s6, s9)
	s2 = uf.Union(s2, s3)
	s2 = uf.Union(s2, s6)

	log.Printf("s2.size = %d", s2.Size)
	log.Printf("2-(3,6-(9)): %v", uf)

	log.Printf("consistent? %v", uf.Consistent())

	// find
	r3 := ref.Find(s3)
	log.Printf("root(3) = %v", r3)
	r6 := ref.Find(s6)
	log.Printf("root(6) = %v", r6)
	r9 := ref.Find(s9)
	log.Printf("root(9) = %v", r9)

	// ufdsref, after path compression
	log.Printf("ufdsref: %v", ref)

	return ref

}

func TestUFDSRef2() *UFDSRef {
	ref := NewUFDSRef("c-5")
	// makeset
	ref.MakeSet(5)
	ref.MakeSet(8)
	ref.MakeSet(10)
	ref.MakeSet(11)
	ref.MakeSet(12)

	log.Printf("all singletons: %v", ref)

	uf := ref.Set

	s11, i11 := ref.GetSingleton(11)
	log.Printf("s6: %v, %v", s11, uf.NodeAt(i11))

	s5, _ := ref.GetSingleton(5)
	s8, _ := ref.GetSingleton(8)
	s10, _ := ref.GetSingleton(10)
	s12, _ := ref.GetSingleton(12)

	// union
	s8 = uf.Union(s8, s11)
	s5 = uf.Union(s5, s10)
	s5 = uf.Union(s5, s8)
	s5 = uf.Union(s5, s12)

	log.Printf("s5.size = %d", s5.Size)
	log.Printf("5-(8-(11),10-(12)): %v", uf)

	log.Printf("consistent? %v", uf.Consistent())

	// find
	r8 := uf.Find(s8)
	log.Printf("root(8) = %v", r8)
	r11 := uf.Find(s11)
	log.Printf("root(11) = %v", r11)
	r12 := uf.Find(s12)
	log.Printf("root(12) = %v", r12)

	// after path compression
	log.Printf("5-(8-(11),10-(12)): %v", uf)

	// ufdsref
	// log.Printf("ufdsref: %v", ufRef)

	return ref

}

func MegaUFDSTest() {
	rf1 := TestUFDSRef1()
	rf2 := TestUFDSRef2()

	// union
	rf1 = rf1.Union(rf2)
	log.Printf("ufdsref: %v", rf1)

	// after mega union
	log.Printf("consistent? %v", rf1.Consistent())

	// find and path compression
	s3, i3 := rf1.GetSingleton(3)
	log.Printf("node %v at index %d", s3, i3)
	s9, i9 := rf1.GetSingleton(9)
	log.Printf("node %v at index %d", s9, i9)
	s10, i10 := rf1.GetSingleton(10)
	log.Printf("node %v at index %d", s10, i10)
	s11, i11 := rf1.GetSingleton(11)
	log.Printf("node %v at index %d", s11, i11)

	uf := rf1.Set

	// find
	r3 := uf.Find(s3)
	log.Printf("root(3) = %v", r3)
	r9 := uf.Find(s9)
	log.Printf("root(9) = %v", r9)
	r10 := uf.Find(s10)
	log.Printf("root(10) = %v", r10)
	r11 := uf.Find(s11)
	log.Printf("root(11) = %v", r11)

	// after path compression
	log.Printf("ufdsref: %v", rf1)

}
