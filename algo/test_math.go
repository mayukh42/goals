package algo

import "log"

func TestDigitsNumber() {
	xs := []int{1, 231, 4567, 895643, 10101, 123542}
	for _, x := range xs {
		ys := Digits(x)
		log.Printf("digits of %d: %v", x, ys)

		y := Number(ys)
		log.Printf("number from %v: %d", ys, y)
	}
}

func TestNextHigherInArr() {
	xs := []int{1}
	for _, x := range xs {
		ys := Digits(x)
		d := byte(1)
		m2i := NextHigherInArr(ys, d)
		log.Printf("next-higher(%d): index = %d, %d", d, m2i, ys[m2i])
	}
}

func TestNextHigherNum() {
	xs := []int{123542}
	for _, x := range xs {
		n := NextHigherNum(x)
		log.Printf("next-higher(%d): %d", x, n)
	}
}
