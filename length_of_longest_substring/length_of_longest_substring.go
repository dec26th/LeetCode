package length_of_longest_substring

import "math"

func lengthOfLongestSubstring(s string) int {
	lenOfString := len(s)
	distinctSet := make(map[byte]int)
	lenOfSubString := 0

	count := -1
	for i := 0; i < lenOfString; i++ {

		if i != 0 {
			delete(distinctSet, s[i - 1])
		}

		for count + 1 < lenOfString && distinctSet[s[count + 1]] == 0  {
			distinctSet[s[count + 1]] ++
			count ++
		}

		lenOfSubString = int(math.Max(float64(lenOfSubString), float64(count - i + 1)))

	}
	return lenOfSubString
}
