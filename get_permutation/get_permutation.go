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

		if num < kth && num > 0 {
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