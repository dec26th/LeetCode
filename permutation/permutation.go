package permutation

var strMap map[string]int

func permutation(s string) []string {
	strList := make([]string, 0)
	strMap = make(map[string]int)
	dfs([]byte(s), 0, &strList)
	return strList
}

func dfs(byteList []byte, start int, strs *[]string) {
	if start == len(byteList) {
		if _, ok := strMap[string(byteList)]; !ok {
			*strs = append(*strs, string(byteList))
			strMap[string(byteList)] = 1
		}

	} else {
		for i := start; i < len(byteList); i++ {
			if i != start {
				byteList[start], byteList[i] = byteList[i], byteList[start]
			}

			dfs(byteList, start+1, strs)

			if i != start {
				byteList[start], byteList[i] = byteList[i], byteList[start]
			}
		}
	}
}
