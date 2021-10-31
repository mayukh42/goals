package osal

import "log"

func TestBitStr() {
	xs := []int{1, 2, 3, 4, 5, 6, 7}
	for _, x := range xs {
		log.Printf("bitstr(%d) = %s", x, BitStr(x))
	}
}

func TestByteStream() {
	xs := []int{4, 7, 10, 64, 100, 127, 161, 242}
	bs := Int8sToByteStream(xs)
	log.Printf("ByteStream of 8-bit ints: %v", bs)
}

func TestLRURefMatrixSet() {
	bs := NewLRURefMatrix()
	log.Printf("ByteStream of 8-bit ints: %v", bs)

	ids := []int{3, 1, 0, 2, 0, 3, 2}
	for _, i := range ids {
		bs, _ = LRURefMatrixSet(bs, i)
		log.Printf("ByteStream set at %d: %v", i, bs)
	}
}
