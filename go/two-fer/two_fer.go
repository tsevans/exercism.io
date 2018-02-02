// Package twofer is an exercism ecercis to print strings
package twofer

// ShareWith prints a given name in a phrase related to twofers
func ShareWith(s string) string {
	if s == "" {
		return "One for you, one for me."
	}

	return "One for " + s + ", one for me."
}
