package osal

import (
	"fmt"
	"strings"
)

func BitStr(n int) string {
	return fmt.Sprintf("%b", n)
}

type ByteStream struct {
	sizeBytes uint
	sizeBits  uint
	Bytes     []byte
}

func NewByteStream(size int) *ByteStream {
	return &ByteStream{
		sizeBytes: uint(size),
		sizeBits:  uint(size) * 8,
		Bytes:     make([]byte, size),
	}
}

func Int8sToByteStream(ns []int) *ByteStream {
	size := len(ns)
	bs := NewByteStream(size)
	for i := range bs.Bytes {
		if ns[i] > 255 {
			continue
		}
		bs.Bytes[i] = byte(ns[i])
	}
	return bs
}

func (b *ByteStream) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%d:%d] \n\t", b.sizeBytes, b.sizeBits))
	for i := range b.Bytes {
		sb.WriteString(PadByte(b.Bytes[i]))
		sb.WriteString("\n\t")
	}
	sb.WriteString("\n")
	return sb.String()
}

func PadByte(b byte) string {
	bs := fmt.Sprintf("%b", b)
	if len(bs) < 8 {
		pd := ""
		for i := 8 - len(bs); i > 0; i-- {
			pd += "0"
		}
		return fmt.Sprintf("%s%s", pd, bs)
	}
	return bs
}

func NewLRURefMatrix() *ByteStream {
	bs := NewByteStream(8)
	for i := range bs.Bytes {
		p := 7 - i
		bs.Bytes[i] = byte((1 << p) - 1)
	}
	return bs
}

func LRURefMatrixSet(bs *ByteStream, n int) (*ByteStream, error) {
	if bs.sizeBytes > 8 || n > 7 {
		return nil, fmt.Errorf("bytestream has wrong size or index is invalid for LRU ref mat")
	}

	// set all bits to 1 in nth row
	bs.Bytes[n] = byte(255)

	// set all bits to 0 in nth col
	mb := int8(1 << (7 - n))
	for i := range bs.Bytes {
		bi := int8(bs.Bytes[i])
		oned := bi | mb
		zeroed := oned ^ mb
		bs.Bytes[i] = byte(zeroed)
	}

	return bs, nil
}

func LRUEvictCandidate(bs *ByteStream) int {
	for i := range bs.Bytes {
		if int(bs.Bytes[i]) == 0 {
			return i
		}
	}
	return -1
}

func LRUEvict(bs *ByteStream) (*ByteStream, error) {
	n := LRUEvictCandidate(bs)
	if n > 0 && n < 8 {
		bs.Bytes[n] = byte(255)
	} else {
		return nil, fmt.Errorf("wrong evict candidate found")
	}

	return bs, nil
}
