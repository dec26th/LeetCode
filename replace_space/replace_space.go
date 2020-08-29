package replace_space

import "strings"

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
