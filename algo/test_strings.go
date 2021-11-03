package algo

import "log"

func TestGetIndex() {
	bts := []byte{'a', 'b', '0', 'm', 'z'}
	for _, b := range bts {
		log.Printf("char-index(%v) = %d", b, getIndex(b))
	}
}

func TestPadString() {
	ss := []string{"alice", "bob", "charles", "david", "eve", "frank", "geeta"}
	for _, s := range ss {
		log.Printf("pad(%s) = %s", s, pad(s, 7))
	}
}

func TestBinSort() {
	xs := []string{"tomason", "geeta", "horace", "giselle", "ada", "gabrielle", "bryan", "abba", "henry", "gabriella", "bob", "tomasok", "alice", "charlie", "hank"}

	ys := FauxLinearSort(xs)
	log.Printf("binsorted: \n\t %v", ys)
}
