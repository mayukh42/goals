package algo

import (
	"log"

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
		log.Printf("x = %b, x & 1 = %b, 2 ** %d = %d", x, x&1, i, p2i)
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
