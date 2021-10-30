package osal

import (
	// l "github.com/mayukh42/goals/list"
	"log"

	u "github.com/mayukh42/goals/utils"
)

func LRUTest() {
	lru := NewLRUCache(4)

	// add 1,2,3
	lru.Set(1, "alice")
	lru.Set(2, "bob")
	lru.Set(3, "charlie")
	// snapshot 1,2,3
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	var (
		v   u.Any
		err error
	)
	// get 2
	v, err = lru.Get(2)
	if err != nil {
		log.Printf("error in get(2): %v", err)
	}
	log.Printf("get(2): %v", v)
	// snapshot 1,3,2
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// get 1
	v, err = lru.Get(1)
	if err != nil {
		log.Printf("error in get(1): %v", err)
	}
	log.Printf("get(1): %v", v)
	// snapshot 3,2,1
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// add 4
	lru.Set(4, "david")
	// snapshot 3,2,1,4
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// add 5
	lru.Set(5, "eve")
	// snapshot 2,1,4,5 (3 evicted)
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// get 4
	v, err = lru.Get(4)
	if err != nil {
		log.Printf("error in get(4): %v", err)
	}
	log.Printf("get(4): %v", v)
	// snapshot 2,1,5,4
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// get 4
	v, err = lru.Get(4)
	if err != nil {
		log.Printf("error in get(4): %v", err)
	}
	log.Printf("get(4): %v", v)
	// snapshot 2,1,5,4
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// get 1
	v, err = lru.Get(1)
	if err != nil {
		log.Printf("error in get(1): %v", err)
	}
	log.Printf("get(1): %v", v)
	// snapshot 2,5,4,1
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// get 2
	v, err = lru.Get(2)
	if err != nil {
		log.Printf("error in get(2): %v", err)
	}
	log.Printf("get(2): %v", v)
	// snapshot 5,4,1,2
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// add 6,7,8
	lru.Set(6, "frank")
	lru.Set(7, "giselle")
	lru.Set(8, "henry")
	// snapshot 2,6,7,8 (evicts 5,4,1)
	log.Printf("LRU Snapshot: %v\n", lru.Snapshot())

	// ...

}
