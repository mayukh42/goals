package osal

import (
	"errors"
	"fmt"
	"strings"

	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

/** LRU Cache implementation
 * double ll + hashtable
 * 		invariant: LRU = head, MRU = tail
 * double ll + reference matrix
 */

type LRUCache struct {
	capacity int
	elements *l.DList
	keyMap   map[u.Any]*l.DNode
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		capacity: size,
		elements: l.NewDList(),
		keyMap:   make(map[u.Any]*l.DNode),
	}
}

func (lru *LRUCache) getKeyFromNode(dn *l.DNode) (u.Any, error) {
	if dn.Value == nil {
		return nil, errors.New("invalid node")
	}
	if kv, ok := dn.Value.(*u.Tuple); ok {
		return kv.First, nil
	}
	return nil, errors.New("key not found")
}

func (lru *LRUCache) getValFromNode(dn *l.DNode) (u.Any, error) {
	if dn.Value == nil {
		return nil, errors.New("invalid node")
	}
	if kv, ok := dn.Value.(*u.Tuple); ok {
		return kv.Second, nil
	}
	return nil, errors.New("value not found")
}

func (lru *LRUCache) add(key u.Any, en *l.DNode) error {
	// invariant fn keyMap
	lru.elements = lru.elements.Suffix(en)
	lru.keyMap[key] = en
	return nil
}

func (lru *LRUCache) Set(key, val u.Any) error {
	if lru.elements.GetSize() == lru.capacity {
		// evict LRU element, size--, map--
		nr := lru.elements.RemoveHead()
		nk, err := lru.getKeyFromNode(nr)
		if err == nil {
			// double node ==> value == (key, val); hashmap invariant
			delete(lru.keyMap, nk)
		}
	}

	// at this point, adding to tail is guranteed to be within capacity
	kv := u.NewTuple(key, val)
	en := l.NewDNode(kv)
	return lru.add(key, en)
}

/** Get()
 * the dlist implementation removes node from its position
 * 	and appends to the tail, maintaining invariant of LRU = head, MRU = tail
 */
func (lru *LRUCache) Get(key u.Any) (u.Any, error) {
	// LRU invariant: if element found, move to tail
	var (
		en *l.DNode
		ok bool
	)
	if en, ok = lru.keyMap[key]; !ok {
		return nil, errors.New("key not found")
	}

	// optimize for tail element
	if en.Equals(lru.elements.Tail) {
		// it is already at tail [TODO: update timestamp]
		return lru.getValFromNode(en)
	}

	// remove this node and add to tail
	en = lru.elements.RemoveNode(en)
	lru.add(key, en)

	return lru.getValFromNode(en)
}

func (lru *LRUCache) Snapshot() string {
	var sb strings.Builder
	sb.WriteString(
		fmt.Sprintf("LRU-DL[%v] (%v) \n\t", lru.capacity, lru.elements.Size),
	)
	sb.WriteString(lru.elements.String())
	sb.WriteString("\n\t")
	sb.WriteString(
		fmt.Sprintf("keyMap: %v \n", lru.keyMap),
	)

	return sb.String()
}
