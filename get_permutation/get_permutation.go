package get_permutation

import "strconv"

var (
	totalNum int
	kth int
	mark []bool
	result string
)

func getPermutation(n int, k int) string {
	totalNum, kth = n, k
	if n == k && n == 1 {
		return "1"
	}
	mark = make([]bool, n + 1)
	backTrack( 0)
	return result
}

func backTrack(temp int) {
	if temp == totalNum {
		return
	}
	num := stageMultiply(totalNum - temp - 1)
	for i := 1; i <= totalNum; i++ {
		if mark[i] {
			continue
		}

		if num < kth {
			kth -= num
			continue
		}
		result += strconv.Itoa(i)
		mark[i] = true
		temp++
		backTrack(temp)
		return
	}
}

func stageMultiply(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func getPermutationTwo(n, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i - 1] * i
	}
	k--

	result := ""
	valid := make([]int, n + 1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}

	for i := 1; i <= n; i++ {
		order := k / factorial[n - i] + 1

		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				result += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n - i]
	}

	return result
}