package algo

import (
	"log"

	l "github.com/mayukh42/goals/list"
	u "github.com/mayukh42/goals/utils"
)

/** x^2 + y^2 < n
 */
func PointsInCircle(n int) []u.Tuple {
	// array of max y for each x
	yMaxs := make([]u.Tuple, 0)
	var x, y int
	// x = 0
	for x*x+y*y < n {
		// O(/n)
		y += 1
	}
	y -= 1
	yMaxs = append(yMaxs, *u.NewTuple(x, y))
	// x = 0, y = yMax
	x += 1
	for y >= 0 {
		for x*x+y*y < n {
			yMaxs = append(yMaxs, *u.NewTuple(x, y))
			// O(/n)
			x += 1
		}
		// O(/n)
		y -= 1
	}

	res := make([]u.Tuple, 0)
	log.Printf("max (x, y) inside circle of radius sqrt(%d): %v", n, yMaxs)
	for _, ymx := range yMaxs {
		var (
			ym int
			x  int
			ok bool
		)
		if ym, ok = ymx.Second.(int); !ok {
			continue
		}
		for j := 0; j <= ym; j++ {
			if x, ok = ymx.First.(int); !ok {
				continue
			}
			res = append(res, *u.NewTuple(x, j))
		}
	}

	return res
}

// n = power of 2
func power2(a, n int) int {
	if n == 0 {
		return 1
	}
	p, x := a, n
	for x > 1 {
		p *= p
		x = x >> 1
	}
	return p
}

func Power(a, n int) int {
	if a <= 1 {
		return a
	}
	if n == 0 {
		return 1
	}

	x, i, s := n, 0, 1
	p2i := 1
	for x > 0 {
		// check if rightmost bit is 1
		// log.Printf("x = %b, x & 1 = %b, 2 ** %d = %d", x, x&1, i, p2i)
		if x&1 == 1 {
			p := power2(a, p2i)
			// log.Printf("p = %d", p)
			s *= p
		}

		// bit pos
		i++
		// decimal value of bit pos
		p2i = p2i << 1
		x = x >> 1
	}

	return s
}

func Divide(v, s int) (int, int) {
	i, p := 1, s
	if p > v {
		return 0, v
	}
	if p == v {
		return 1, 0
	}

	// loop until lowest multiple which is power of 2
	for p < v {
		p = p << 1
		if p > v {
			// undo
			p = p >> 1
			break
		}
		i = i << 1
	}

	dv := v - p
	for dv >= s {
		i++
		dv -= s
	}

	return i, dv
}

func NewtonSqrt(n int) int {
	p := 1
	for p < n {
		p = p << 1
		if p >= n {
			// undo
			p = p >> 1
			break
		}
	}

	log.Printf("[sqrt(%d)] p = %d", n, p)

	// p is first approx, max power of 2 that is < n
	a, b := p, n/p
	i := 0
	for u.AbsDiff(uint(a), uint(b)) > 1 {
		a = a + (b-a)/2
		b = n / a
		i++
		// log.Printf("[sqrt(%d)] [%d] a = %d, b = %d", n, i, a, b)
	}

	if a*a > n {
		return b
	}
	return a
}

/** a little more complicated than needs to be
 * but uses collection reverse
 */
func Digits(n int) []byte {
	ds := make(u.List, 0)
	for n > 0 {
		ds = append(ds, byte(n%10))
		n = n / 10
	}
	drs := l.ReverseArr(ds)
	// log.Printf("digits of %d reversed: %v", n, drs)

	dgts := make([]byte, len(drs))
	for i := range drs {
		d, ok := drs[i].(byte)
		if ok {
			dgts[i] = d
		}
	}
	return dgts
}

func Number(digits []byte) int {
	n := 0
	for i := len(digits) - 1; i >= 0; i-- {
		p := len(digits) - i - 1
		if digits[i] > 0 {
			n += Power(10, p) * int(digits[i])
		}

	}
	return n
}

// return index of next higher num than d, in digits array
func NextHigherInArr(digits []byte, d byte) int {
	if len(digits) < 1 {
		return 0
	}
	m2 := digits[0]
	m2i := 0
	for i := 1; i < len(digits); i++ {
		if digits[i] > d {
			if digits[i] < m2 {
				m2 = digits[i]
				m2i = i
			}
		}
	}

	return m2i
}

func NextHigherNum(num int) int {
	ds := Digits(num)
	b := -1
	n := len(ds)
	// find i at the left boundary of non-increasing sequence from end
	for i := n - 1; i > 0; i-- {
		if ds[i-1] < ds[i] {
			b = i - 1
			break
		}
	}

	// split into 3 areas - [0..b-1], [b], [b+1..n-1]
	// first := ds[:b]
	last := ds[b+1:]

	// search for lowest num > ds[b] in [b+1..n-1]
	m := NextHigherInArr(last, ds[b])
	// correct pos (+1 for index starting with 0)
	m += b + 1

	log.Printf("n, b, m: %d, %d, %d", n, b, m)
	// swap m2i with b
	tmp := ds[m]
	ds[m] = ds[b]
	ds[b] = tmp

	// reverse the part [b+1..n-1], i.e. last
	// e.g. [0...2][3,4,5,6,7,8,9]: swap equidistant elements from either side
	for i := b + 1; i < b+((n-b)>>1); i++ {
		tmp := ds[i]
		ds[i] = ds[n-(i-b)]
		ds[n-(i-b)] = tmp
	}

	// construct digits
	return Number(ds)
}

// O(n) for n squares!
func SquarePro(n int) []int {
	sqrs := make([]int, n+1)
	k, k2 := 0, 0
	// invariant
	sqrs[k] = k2

	for k < n {
		k++
		/** once k = k+1, k2 = (k-1)^2
		 * 	(since k2 still has (k-1)^2 stored) = k^2 -2k + 1
		 * 	so we add 2k - 1 to k2, so that it becomes k^2 which we want
		 */
		k2 = k2 + k + k - 1
		sqrs[k] = k2
	}
	return sqrs
}

func GCD(a, b int) int {
	// assume a > b
	if b == 0 || a == b {
		return a
	}
	if b == 1 {
		return b
	}

	if a > b {
		return GCD(b, a-b)
	} else {
		return GCD(a, b-a)
	}
}
