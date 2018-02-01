// Package acronym is an exercism exercise to convert phrases to acronyms
package acronym

import "strings"

// Abbreviate takes a phrase and returns its acronym representation
func Abbreviate(s string) string {

	words := strings.FieldsFunc(s, Split)
	var acronym string

	for _, w := range words {
		l := w[:1]
		acronym += strings.ToUpper(l)
	}

	return acronym
}

func Split(r rune) bool {
	return r == ' ' || r == '-'
}
