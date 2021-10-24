package algo

import (
	"log"

	u "github.com/mayukh42/goals/utils"
)

func TestBinarySearch() {
	xs := u.List{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ns := []int{0, 1, 6, 3, 9, 10, 100}

	for _, x := range ns {
		found := BinarySearch(xs, x, u.CompareInt, 0, len(xs)-1)
		log.Printf("%d is found in %v? %v", x, xs, found)
	}
}
