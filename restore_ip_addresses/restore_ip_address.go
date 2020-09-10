package restore_ip_addresses

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	lenOfS := len(s)
	if lenOfS < 4 || lenOfS > 12 {
		return []string{}
	}
	result := new([]string)
	tempResult := *new([]string)

	dfs(1, []byte(s), result, tempResult)
	return *result
}

func dfs(layer int, numList []byte, result *[]string, tempResult []string){
	if layer > 4 {
		return
	}

	lenOfNum := len(numList)
	if lenOfNum > 3 * (5 - layer) {  // 剩余过多
		return
	}

	if layer == 4 {
		num, _ := strconv.Atoi(string(numList))
		if num <= 255 {
			tempResult = append(tempResult, string(numList))
			if ! judgeValid(tempResult) {
				return
			}
			*result = append(*result, strings.Join(tempResult, "."))
			return
		}
	}
	for i := 1; i < lenOfNum; i ++ {
		if i < 4 {
			num := string(numList[:i])
			if intNum, _ := strconv.Atoi(num); intNum <= 255 {
				temp := tempResult[:]
				temp = append(temp, num)
				dfs(layer + 1, numList[i:], result, temp)
			}
		}
	}
}

func judgeValid(tempResult []string) bool {
	for _, num := range tempResult {
		if len(num) != 1 && num[0] - '0' == 0 {
			return false
		}
	}
	return true
}
