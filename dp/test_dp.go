package dp

import "fmt"

func TestLevenshtein() {
	w1 := "flaw"
	w2 := "lawn"

	lwd := Levenshtein(w1, w2)
	fmt.Printf("Levenshtein(%s, %s) = %d\n", w1, w2, lwd)
}

func TestLCSeq() {
	w1 := "AGGTAB"
	w2 := "GXTXAYB"

	lwd := LCSeq(w1, w2)
	fmt.Printf("LCSeq(%s, %s) = %s\n", w1, w2, lwd)
}

func TestMaximalSum() {
	xs := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := MaximalSum(xs)
	fmt.Printf("MaximalSum(%v) = %d\n", xs, maxSum)
}
