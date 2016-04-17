// Package validation provides a set of functions to validate input
package validation

import (
	"fmt"
	"regexp"
)
// IsAlpha determines if the passed string is all chars
func IsAlpha(input string) {
	rxAlpha := regexp.MustCompile("^[a-zA-Z\\s]+$")
	if regexpMatches(input, rxAlpha) == false {
		panic(fmt.Sprintf("%v needs to be all alpha chars", input))
	}
}
// IsNumeric determines if the passed string is all digits
func IsNumeric(input string) {
	rxNumeric := regexp.MustCompile("^[0-9]+$")
	if regexpMatches(input, rxNumeric) == false {
		panic(fmt.Sprintf("%v needs to be all alphanumeric chars", input))
	}
}

func regexpMatches(input string, rx *regexp.Regexp) bool {
	if len(input) == 0 {
		return true
	}
	return rx.MatchString(input)
}
