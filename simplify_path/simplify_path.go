package simplify_path

import "strings"

func simplifyPath(path string) string {
	splitList := strings.Split(path, "/")
	result := *new([]string)

	for _, char := range splitList {
		switch char {
		case "":
			break
		case ".":
			break
		case "..":
			if l := len(result); l > 0 {
				result = result[:l-1]
			}
		default:
			result = append(result, char)

		}
	}

	return "/" + strings.Join(result, "/")
}
