package validator

import (
	"fmt"
	"net/http"
	"regexp"
)

// New returns a new handler.
func New(next http.Handler) *Handler {
	return &Handler{next}
}

// Handler encapsulates a boffle.
type Handler struct {
	next http.Handler
}

var rePattern = regexp.MustCompile("^/(loaderio-[0-9a-f]+).*")

func (l *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	matches := rePattern.FindStringSubmatch(r.URL.Path)
	if matches != nil {
		w.WriteHeader(200)
		fmt.Fprintln(w, matches[1])
		return
	}
	l.next.ServeHTTP(w, r)
}
