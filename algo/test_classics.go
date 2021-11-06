package algo

import (
	"log"

	u "github.com/mayukh42/goals/utils"
)

func testBinarySearch() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ns := []int{0, 1, 6, 3, 9, 10, 100}

	for _, x := range ns {
		found := BinarySearch(xs, x, u.CompareInt, 0, len(xs)-1)
		log.Printf("%d is found in %v? %v", x, xs, found)
	}
}

func testCountingSort() {
	xs := []uint8{64, 83, 7, 32, 32, 41, 7, 32, 89, 121, 8, 222, 251, 161, 7, 171, 8, 161, 41, 251, 161}
	ys := CountingSort(xs)
	log.Printf("c-sort(%d): \n\t %v", xs, ys)
}

func testRadixSort() {
	xs := []uint8{97, 87, 17, 43, 45, 8, 7, 91, 63, 72, 16, 79}
	cs := RadixSort(xs, 2)
	log.Printf("radix(%v): \n\t %v", xs, cs)
}

func testPartition() {
	xs := []int{64, 83, 7, 32, 92, 41, 7, 32, 89, 121, 8, 222, 251, 161, 7, 171, 8, 161, 43, 251, 161}
	log.Printf("original: \n\t %v", xs)
	pi := Partition(xs, 0, len(xs))
	log.Printf("partitioned at %d: \n\t %v", pi, xs)

	xs = []int{42}
	log.Printf("original: \n\t %v", xs)
	pi = Partition(xs, 0, len(xs))
	log.Printf("partitioned at %d: \n\t %v", pi, xs)

	xs = []int{42, 28}
	log.Printf("original: \n\t %v", xs)
	pi = Partition(xs, 0, len(xs))
	log.Printf("partitioned at %d: \n\t %v", pi, xs)

	xs = []int{251, 251, 222, 222, 201}
	log.Printf("original: \n\t %v", xs)
	pi = Partition(xs, 0, len(xs))
	log.Printf("partitioned at %d: \n\t %v", pi, xs)
}

func testQuicksort() {
	xs := []int{64, 83, 7, 32, 92, 41, 7, 32, 89, 121, 8, 222, 251, 161, 7, 171, 8, 161, 43, 251, 161}
	log.Printf("original: \n\t %v", xs)
	Quicksort(xs, 0, len(xs))
	log.Printf("partitioned: \n\t %v", xs)

	xs = []int{8, 8, 8, 8, 8, 8}
	log.Printf("original: \n\t %v", xs)
	Quicksort(xs, 0, len(xs))
	log.Printf("partitioned: \n\t %v", xs)
}
