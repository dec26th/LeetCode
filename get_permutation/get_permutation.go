package get_permutation

import "strconv"

func getPermutation(n int, k int) string {
	return  backTrack(n, k, 1, "")
}

func backTrack(n, k, layer int, result string) string {
	num := stageMultiply(n - layer)

	if num > k {
		layer ++
		result += strconv.Itoa(layer)
		backTrack(n, k, layer, result)
	}

	return  ""

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