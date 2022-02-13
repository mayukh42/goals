package dp

import (
	"fmt"
	"log"
)

func testLevenshtein() {
	w1 := "flaw"
	w2 := "lawn"

	lwd := Levenshtein(w1, w2)
	fmt.Printf("Levenshtein(%s, %s) = %d\n", w1, w2, lwd)
}

func testLCSeq() {
	w1 := "AGGTAB"
	w2 := "GXTXAYB"

	lwd := LCSeq(w1, w2)
	fmt.Printf("LCSeq(%s, %s) = %s\n", w1, w2, lwd)
}

func testMaximalSum() {
	xss := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{31, -41, 59, 26, -53, 58, 97, -93, -23, 84},
	}
	for _, xs := range xss {
		maxSum := MaximalSum(xs)
		fmt.Printf("MaximalSum(%v) = %d\n", xs, maxSum)
	}
}

func testClosestP() {
	str := "googlee"
	ClosestPalindrome(str)
}

func testIsPalindromeCandidate() {
	ss := []string{"aeiou", "googlee", "goog"}
	for _, s := range ss {
		pc := IsPalindromeCandidate(s)
		log.Printf("can %s be palindrome? %v", s, pc)
	}
}

func TestDP() {
	// testLevenshtein()
	// testLCSeq()
	// testMaximalSum()
	// testClosestP()
	testIsPalindromeCandidate()
}

func TestDPAll() {
	testLevenshtein()
	testLCSeq()
	testMaximalSum()
	testClosestP()
	testIsPalindromeCandidate()
}
