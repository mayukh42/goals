package algo

import (
	"log"
)

func testSpiralWalk() {
	xy := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}
	walk := SpiralWalk(xy, 4, 5)
	log.Printf("spiral walk: %v", walk)
}

func testMatAt() {
	cs := [][]int{
		{1, 2},
		{0, 1},
		{2, 3},
		{3, 3},
		{4, 2},
	}
	xy := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}
	for _, c := range cs {
		i, j := c[0], c[1]
		e, _ := MatAt(xy, 5, 4, i, j)
		log.Printf("xy(%d,%d) = %d", i, j, e)
	}
}

func testSaddlebackSearchL() {
	xy := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 15, 15, 16},
		{17, 18, 19, 20},
	}
	es := []int{15, 1, 12, 20, 25}
	for _, e := range es {
		i, j := SaddlebackSearchL(xy, 5, 4, e)
		log.Printf("%d found at coords (%d,%d)", e, i, j)
	}
}

func testSaddlebackSearch() {
	xy := [][]int{
		{2, 2, 3, 5},
		{3, 4, 5, 6},
		{3, 5, 6, 8},
		{3, 6, 7, 9},
	}
	es := []int{6, 5, 9, 4, 2, 10}
	for _, e := range es {
		i, j := SaddlebackSearch(xy, 4, 4, e)
		log.Printf("%d found at coords (%d,%d)", e, i, j)
	}
}
