package collections

import (
	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

type Heap struct {
	Array      l.List
	Size       int
	Count      int
	Comparator u.Comparator
}

func LowerInt(u1, u2 u.Any) bool {
	if u1i, ok := u1.(int); ok {
		if u2i, ok := u2.(int); ok {
			return u1i < u2i
		}
	}
	return false
}

func NewHeap(xs l.List, comp u.Comparator) *Heap {
	size := len(xs) + 1
	l := make(l.List, size)
	// shift all elements of xs into a new array w/ index + 1
	for i := 1; i <= size; i++ {
		l[i] = xs[i-1]
	}

	return &Heap{
		Array:      l,
		Size:       size,
		Count:      size,
		Comparator: comp,
	}
}

func (h *Heap) Remove() u.Any {
	if len(h.Array) >= 2 {
		i, j := 1, h.Count+1
		first, last := h.Array[i], h.Array[j]
		h.Array[j] = 0
		h.Array[i] = last
		h.Count--
		h.siftdown(i)
		return first
	}
	return nil
}

func (h *Heap) Insert(e u.Any) error {
	h.Array = append(h.Array, e)
	h.Count++
	h.Size++
	// restore heap invariant
	return h.siftup(h.Count)
}

/** siftup()
 * shifts element at index i up, maintaining heap invariant
 * 	invariant:
 * 		parent is always smaller[minheap] or greater[maxheap]
 * 		than either children
 * 	bias: leftward
 */
func (h *Heap) siftup(i int) error {
	if i <= 1 {
		return nil
	}
	// if xs[i] < xs[i/2] then swap
	if h.Comparator(h.Array[i], h.Array[i/2]) {
		h.swap(i, i/2)
		return h.siftup(i / 2)
	}
	return nil
}

/** siftdown()
 * shifts element at index i down, maintaining heap invariant
 * 	invariant:
 * 		parent is always smaller[minheap] or greater[maxheap]
 * 		than either children
 * 	bias: leftward
 */
func (h *Heap) siftdown(i int) error {
	if i > h.Size {
		return nil
	}
	// if xs[i] < xs[i*2] then swap
	if h.Comparator(h.Array[i], h.Array[i*2]) {
		h.swap(i, i*2)
		return h.siftdown(i * 2)
	}
	return nil
}

func (h *Heap) swap(i, j int) error {
	tmp := h.Array[i]
	h.Array[i] = h.Array[j]
	h.Array[j] = tmp
	return nil
}
