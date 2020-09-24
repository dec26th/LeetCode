package num_ways

func numWays(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}
	way := make([]int, n+1)
	way[1], way[2] = 1, 2
	for i := 3; i < n+1; i++ {
		way[i] = (way[i-1] + way[i-2]) % (1e9 + 7)
	}
	return way[n]
}
