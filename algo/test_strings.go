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
