package longest_common_prefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1{
		return strs[0]
	}
	var tempLongestPrefix string
	for k, str := range strs {
		if k == 0 {
			tempLongestPrefix = str
			continue
		}

		lenOfTempPrefix := len(tempLongestPrefix)
		for k := 0;  k < lenOfTempPrefix; k++ {
			if ! strings.HasPrefix(str, tempLongestPrefix[:k + 1]) {
				if k == 0 {
					return ""
				}
				tempLongestPrefix = tempLongestPrefix[:k]
				break
			}
		}
	}
	if strs[0] == tempLongestPrefix && ! strings.Contains(strs[1], tempLongestPrefix) {
		tempLongestPrefix = ""
	}
	return tempLongestPrefix
}