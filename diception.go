package diception

import (
	"regexp"
	"strings"
)

// OMG is a dumb function that responds if a word starting with `di` is in a french sentence.
func OMG(s string) string {
	re := regexp.MustCompile("(di[a-z]{2,}|dy[a-z]{2,})")

	sm := re.FindString(s)

	if len(sm) > 0 {
		re := regexp.MustCompile("di|dy")
		return strings.Trim(re.ReplaceAllString(sm, ""), " ")
	}

	return ""
}
