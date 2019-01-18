package routes

import (
	"regexp"
	"strings"
	"net/http"
	"io"
)

// OMG is a dumb function that responds if a word starting with `di` is in a french sentence.
func OMG(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("s")
		
	// Create the regex
	re := regexp.MustCompile("(di[a-z]{2,}|dy[a-z]{2,}|Di[a-z]{2,}|d\\'i[a-z]{2,})")

	// Find string
	sm := re.FindString(s)

	// Trim prefix
	sm = strings.TrimPrefix(sm, "di")
	sm = strings.TrimPrefix(sm, "dy")
	sm = strings.TrimPrefix(sm, "Di")
	sm = strings.TrimPrefix(sm, "d'i")

	io.WriteString(w, sm)
}

