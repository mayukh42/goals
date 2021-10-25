package cp

import (
	"log"
	"sync"

	u "github.com/mayukh42/goals/utils"
)

// Thread Safe object
type TSO struct {
	mtx sync.Mutex
	idx int
}

func (s *TSO) Change(v int) {
	s.mtx.Lock()
	s.idx = v
	s.mtx.Unlock()
}

func ConcurrentMinMax(xs u.List, cf u.Comparator) u.Any {
	acc := pCompare(xs, cf)
	if len(acc) < 1 {
		return u.List{}
	}
	return acc[0]
}

func ConcurrentSumProduct(xs u.List, cf u.BinaryInt) u.Any {
	acc := pSumProduct(xs, cf)
	if len(acc) < 1 {
		return u.List{}
	}
	return acc[0]
}

func divideByTwo(xs u.List) u.List {
	lx := len(xs)
	la := 1
	if lx == (lx>>1)<<1 {
		// check if lx = even
		la = lx >> 1
	} else {
		la = (lx >> 1) + 1
	}
	acc := make(u.List, la)
	return acc
}

/** pCompare()
 * compare every 2 elements i, i+1 = j of xs, and put in idx = i/2 of acc
 */
func pCompare(xs u.List, cf u.Comparator) u.List {
	lx := len(xs)
	if lx <= 1 {
		// base case
		return xs
	}

	acc := divideByTwo(xs)
	wg := &sync.WaitGroup{}
	for i := 0; i < lx; i = i + 2 {
		j := i + 1
		log.Printf("spawning sheep %d\n", i)
		wg.Add(1)
		go grCompareAndStore(xs, acc, i, j, cf, wg)
	}

	wg.Wait()
	log.Printf("all sheep returned\n")
	// all done in this round; now acc is the new xs
	return pCompare(acc, cf)
}

/** smug goroutine; could use some sleep?
 */
func grCompareAndStore(xs, acc u.List, i, j int, cf u.Comparator, wg *sync.WaitGroup) {
	defer wg.Done()
	var winner u.Any
	if cf(xs[i], xs[j]) {
		// i wins
		winner = xs[i]
	} else {
		winner = xs[j]
	}
	acc[i>>1] = winner
	log.Printf("winner: %v\n", winner)
}

func pSumProduct(xs u.List, f u.BinaryInt) u.List {
	lx := len(xs)
	if lx <= 1 {
		// base case
		return xs
	}

	acc := divideByTwo(xs)

	wg := &sync.WaitGroup{}
	for i := 0; i < lx; i = i + 2 {
		j := i + 1
		log.Printf("spawning sheep %d\n", i)
		wg.Add(1)
		go grOpAndStore(xs, acc, i, j, f, wg)
	}

	wg.Wait()
	log.Printf("all sheep returned\n")
	// all done in this round; now acc is the new xs
	return pSumProduct(acc, f)
}

func grOpAndStore(xs, acc u.List, i, j int, f u.BinaryInt, wg *sync.WaitGroup) {
	defer wg.Done()
	if a, ok := xs[i].(int); ok {
		if b, ok := xs[j].(int); ok {
			sum := f(a, b)
			acc[i>>1] = sum
			log.Printf("%d + %d = %d\n", a, b, sum)
		}
	}
}
