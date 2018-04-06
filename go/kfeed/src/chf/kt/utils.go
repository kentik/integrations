package kt

import "strings"

func SingleLineElide(s string, maxLen int) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	if len(s) > maxLen && maxLen >= 3 {
		s = s[:maxLen-3] + "..."
	}
	return s
}
