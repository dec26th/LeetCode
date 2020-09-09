package check_inclusion

import "strings"
//todo 排列问题
func checkInclusion(s1 string, s2 string) bool  {
	resultList := permutation(s2)

	for _, posbString := range resultList {
		if strings.Contains(s1, posbString) {
			return true
		}
	}
	return false
}

func permutation(s string) []string {
	resultList := make([]string, 0)
	lenOfString := len(s)

	for i := 0; i < lenOfString; i++ {

	}
	return resultList
}
