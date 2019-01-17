package diception

import (
	"regexp"
	"strings"
)

// OMG is a dumb french function that responds if a word is starting with `di`.
func OMG(s string) string {
	re := regexp.MustCompile("(di\\w+|dy\\w+)")

	sm := re.FindString(s)

	if len(sm) > 0 {
		re := regexp.MustCompile("di|dy")
		return strings.Trim(re.ReplaceAllString(sm, ""), " ")
	}

	return ""
}
