package arr

import "fmt"

// Contains checks the value exists in Array.
func Contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// Join returns a string created by converting each string, separated by the given separator.
func Join(s []string, sep string) string {
	var buf string
	for _, v := range s[:len(s)-1] {
		buf += fmt.Sprintf("%s%s", v, sep)
	}

	buf += s[len(s)-1]
	return buf
}
