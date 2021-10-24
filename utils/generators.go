package utils

func IntList(min, max, diff int) List {
	i := min
	xs := make(List, 0)
	for i < max {
		xs = append(xs, i)
		i += diff
	}

	return xs
}

func StringList(n int) {
	// permutations
}
