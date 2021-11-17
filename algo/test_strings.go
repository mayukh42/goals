package algo

import "log"

func testGetIndex() {
	bts := []byte{'a', 'b', '0', 'm', 'z'}
	for _, b := range bts {
		log.Printf("char-index(%v) = %d", b, getIndex(b))
	}
}

func testPadString() {
	ss := []string{"alice", "bob", "charles", "david", "eve", "frank", "geeta"}
	for _, s := range ss {
		log.Printf("pad(%s) = %s", s, pad(s, 7))
	}
}

func testBinSort() {
	xs := []string{"tomason", "geeta", "horace", "giselle", "ada", "gabrielle", "bryan", "abba", "henry", "gabriella", "bob", "tomasok", "alice", "charlie", "hank"}

	ys := FauxLinearSort(xs)
	log.Printf("binsorted: \n\t %v", ys)
}

func testSuffixArray() {
	ss := []string{"banana"}
	for _, s := range ss {
		sa := NewSuffixArray(s)
		log.Printf("suffix_array(%s): %s", s, sa)
	}
}

/**
 * TODO: test with set
 */
func testPermutations() {
	ss := []string{"aeiou"}
	for _, s := range ss {
		ps := Permutations(s)
		log.Printf("permutations(%s): \n\t %v, \n\t size: %d", s, ps, len(ps))
	}
}

func testSubsets() {
	ss := []string{"aeiou"}
	for _, s := range ss {
		ps := Subsets(s, []string{""})
		log.Printf("subsets(%s): \n\t %v, \n\t size: %d",
			s, ps, len(ps))
	}
}

func testCombinations() {
	ss := []string{"aeiou"}
	for _, s := range ss {
		ps := Combinations(s)
		log.Printf("combinations(%s): \n\t %v, \n\t size: %d", s, ps, len(ps))
	}
}
