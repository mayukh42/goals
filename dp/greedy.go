package dp

import (
	"log"

	"github.com/mayukh42/goals/utils"
)

func ClosestPalindrome(str string) string {
	// first element as char, second as nearestdistance from center
	evenMap := make(map[byte]utils.Tuple)
	oddMap := make(map[byte]utils.Tuple)

	fMap := make(map[byte]int)
	for _, r := range str {
		b := byte(r)
		if n, ok := fMap[b]; ok {
			fMap[b] = n + 1
		} else {
			fMap[b] = 1
		}
	}

	l := len(str)
	// center to end
	for i := (l >> 1) + 1; i < l; i++ {
		d := i - l/2
		b := byte(str[i])
		if n, ok := fMap[b]; ok {
			// even
			if n == (n>>1)<<1 {
				if t, ok := evenMap[b]; ok {
					t2, ok := t.Second.(int)
					if ok {
						if d < t2 {
							// change
							evenMap[b] = *utils.NewTuple(b, d)
						}
					}
				} else {
					evenMap[b] = *utils.NewTuple(b, d)
				}
			} else {
				// odd
				if t, ok := oddMap[b]; ok {
					t2, ok := t.Second.(int)
					if ok {
						if d < t2 {
							// change
							oddMap[b] = *utils.NewTuple(b, d)
						}
					}
				} else {
					oddMap[b] = *utils.NewTuple(b, d)
				}
			}
		}
	}

	// center to start
	for i := (l >> 1); i >= 0; i-- {
		d := l/2 - i
		b := byte(str[i])
		if n, ok := fMap[b]; ok {
			// even
			if n == (n>>1)<<1 {
				if t, ok := evenMap[b]; ok {
					t2, ok := t.Second.(int)
					if ok {
						if d < t2 {
							// change
							evenMap[b] = *utils.NewTuple(b, d)
						}
					}
				} else {
					evenMap[b] = *utils.NewTuple(b, d)
				}
			} else {
				// odd
				if t, ok := oddMap[b]; ok {
					t2, ok := t.Second.(int)
					if ok {
						if d < t2 {
							// change
							oddMap[b] = *utils.NewTuple(b, d)
						}
					}
				} else {
					oddMap[b] = *utils.NewTuple(b, d)
				}
			}
		}
	}

	// evens, odds
	evens := make([]utils.Tuple, 0)
	for _, v := range evenMap {
		evens = append(evens, v)
	}

	odds := make([]utils.Tuple, 0)
	for _, v := range oddMap {
		odds = append(odds, v)
	}

	log.Printf("evens: %v", evens)
	log.Printf("odds: %v", odds)

	return str
}
