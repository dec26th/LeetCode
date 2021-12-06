package length_of_longest_substring

import "math"

func lengthOfLongestSubstring(s string) int {
	lenOfString := len(s)
	distinctSet := make(map[byte]int)
	lenOfSubString := 0

	count := -1
	for i := 0; i < lenOfString; i++ {

		if i != 0 {
			delete(distinctSet, s[i-1])
		}

		for count+1 < lenOfString && distinctSet[s[count+1]] == 0 {
			distinctSet[s[count+1]]++
			count++
		}

		lenOfSubString = int(math.Max(float64(lenOfSubString), float64(count-i+1)))

	}
	return lenOfSubString
}


func lengthOfLongestSubstringTwo(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	record := make(map[uint8]struct{})

	max := 0
	length := 0
	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(record, s[i - 1])

		}
		for j := length + i; j < len(s); j++ {
			if _, ok := record[s[j]]; !ok {
				record[s[j]] = struct{}{}
				length ++
			} else {
				break
			}
		}

		if length > max {
			max = length
		}
		length--
	}

	return max
}
