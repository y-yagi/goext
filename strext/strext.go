package strext

import "regexp"

var blankRe = regexp.MustCompile(`\A[[:space:]]*\z`)

// IsBlank checks the string is empty or not.
func IsBlank(str string) bool {
	return len(str) == 0 || blankRe.MatchString(str)
}
