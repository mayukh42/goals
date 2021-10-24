package list

import (
	"log"

	u "github.com/mayukh42/goals/utils"
)

func CreateAndPrintIntList() {
	xs := List{1, 2, 3, 4, 5}
	ll := CreateFromArray(xs)
	log.Printf("list: %v\n", ll.String())
}

func CreateAndPrintStringList() {
	xs := List{"facebook", "apple", "amazon", "netflix", "google"}
	ll := CreateFromArray(xs)
	log.Printf("list: %v\n", ll.String())
}

func DeleteIntList() {
	xs := List{1, 2, 3, 4, 5}
	ll := CreateFromArray(xs)
	log.Printf("list: %v\n", ll.String())

	DeleteNode(ll)
}

func SplitMergeIntList() {
	xs := List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ll := CreateFromArray(xs)
	log.Printf("list: %v\n", ll.String())

	first, second := SplitIntoLists(ll)
	log.Printf("first list: %v\n", first.String())
	log.Printf("second list: %v\n", second.String())

	ll2 := MergeLists(first, second)
	log.Printf("merged: %v\n", ll2.String())
}

func TestReverse() {
	xs := List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ll := CreateFromArray(xs)
	log.Printf("list: %v\n", ll.String())

	lr := Reverse(ll)
	log.Printf("reversed: %v\n", lr.String())
}

func TestRotate() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	log.Printf("original: %v\n", xs)

	var (
		ys u.List
		k  uint
	)

	k = 4

	ys = RotateInPlace(xs, k)
	log.Printf("rotated: %v\n", ys)

	ys = SwapPrefixSuffix(xs, k)
	log.Printf("rotated: %v\n", ys)

}

func TestDListSuffixed() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dl := CreateDListSuffixed(xs)

	log.Printf("double linked list suffixed: %v\n", dl)
}

func TestDListPrefixed() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dl := CreateDListPrefixed(xs)

	log.Printf("double linked list prefixed: %v\n", dl)
}
