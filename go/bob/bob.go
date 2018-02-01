// Package bob is an exercism exercise to simulate interactions with a lackadaisical teenager named Bob
package bob

import "strings"
import "unicode"

// Hey takes a phrase as input and returns Bob's response to it.
func Hey(remark string) string {

	//Clean up input string
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	var question bool = strings.HasSuffix(remark, "?")
	var yell bool = allUppercase(remark)

	if question && yell {
		return "Calm down, I know what I'm doing!"
	} else if question {
		return "Sure."
	} else if yell {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

// allUppercase checks if all letters in remark are uppercase
func allUppercase(remark string) bool {
	var hasLetter bool = false
	var upcase bool = true
	for _, x := range remark {
		if unicode.IsLetter(x) {
			hasLetter = true
			if !unicode.IsUpper(x) {
				upcase = false
				break
			}
		}
	}

	return hasLetter && upcase
}
