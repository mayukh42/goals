package utils

import "fmt"

const MAX_DIM = 100

func Max(x, y uint) uint {
	if x > y {
		return x
	} else {
		return y
	}
}

func Min(x, y uint) uint {
	if x < y {
		return x
	} else {
		return y
	}
}

func AbsDiff(x, y uint) uint {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func Min3(x, y, z uint) uint {
	return Min(x, Min(y, z))
}

func Print2DMatrix(mat [MAX_DIM][MAX_DIM]int, l1, l2 int) {
	if l1 < 0 {
		l1 = 0
	}
	if l1 > MAX_DIM {
		l1 = MAX_DIM
	}
	if l2 < 0 {
		l2 = 0
	}
	if l2 > MAX_DIM {
		l2 = MAX_DIM
	}

	fmt.Printf("matrix[%d][%d]\n", l2, l1)
	for i := 0; i < l1; i++ {
		fmt.Printf("[ ")
		for j := 0; j < l2; j++ {
			fmt.Printf("%v, ", mat[i][j])
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("\n")
}

func Swap(xs List, a, b uint) List {
	n := uint(len(xs))
	if a >= n || b >= n {
		return xs
	}
	tmp := xs[a]
	xs[a] = xs[b]
	xs[b] = tmp

	return xs
}
