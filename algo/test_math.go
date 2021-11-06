package algo

import (
	"log"
)

func testDigitsNumber() {
	xs := []int{1, 231, 4567, 895643, 10101, 123542}
	for _, x := range xs {
		ys := Digits(x)
		log.Printf("digits of %d: %v", x, ys)

		y := Number(ys)
		log.Printf("number from %v: %d", ys, y)
	}
}

func testNextHigherInArr() {
	xs := []int{1}
	for _, x := range xs {
		ys := Digits(x)
		d := byte(1)
		m2i := NextHigherInArr(ys, d)
		log.Printf("next-higher(%d): index = %d, %d", d, m2i, ys[m2i])
	}
}

func testNextHigherNum() {
	xs := []int{123542}
	for _, x := range xs {
		n := NextHigherNum(x)
		log.Printf("next-higher(%d): %d", x, n)
	}
}

func testSquarePro() {
	n := 10
	sqrs := SquarePro(n)
	log.Printf("squares[0..%d] = %v", n, sqrs)
}

func testGCD() {
	xys := [][2]int{
		{18, 12},
		{27, 24},
		{30, 31},
		{405, 360},
		{1024, 64},
		{198, 1331},
	}
	for _, xy := range xys {
		a, b := xy[0], xy[1]
		log.Printf("gcd(%d, %d) = %d", a, b, GCD(a, b))
	}
}
