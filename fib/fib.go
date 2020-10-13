package fib

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	var now int
	pp, p := 0, 1

	for i := 2; i <= N; i++ {
		now = (p + pp) % (1e9 + 7)
		pp = p
		p = now
	}
	return now
}
