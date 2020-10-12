package my_sqrt

func mySqrt(x int) int {
	if x == 1 {
		return x
	}
	biSearch := x / 2
	left, right := 0, x
	for {
		if sqrt(biSearch) <= x && sqrt(biSearch+1) > x {
			return biSearch
		}
		if sqrt(biSearch) < x {
			left = biSearch + 1
			biSearch = (right + biSearch) / 2
			continue
		}

		if sqrt(biSearch) > x {
			right = biSearch
			biSearch = (left + biSearch) / 2
			continue
		}

	}
}

func sqrt(x int) int {
	return x * x
}
