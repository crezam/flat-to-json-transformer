// Package processing provides a set of functions to transform input
package processing

import (
	"./validation"
	"strconv"
	"strings"
	"time"
)

// ProcessData transforms from mm-dd-yyyy to mm/dd/yyyy format
func ProcessDate(dateIn string) string {
	dateOut, err := time.Parse("01-02-2006", dateIn)
	// relay validation performed by time package
	if err != nil {
		panic(err)
	}
	// expected format in output
	return dateOut.Format("01/02/2006")
}

// ProcessCard currently does not change card number
func ProcessCard(card string) string {
	validation.IsNumeric(card)
	//manipulation would go here
	return card
}

// ProcessName trims any pre/tail space
func ProcessName(name string) string {
	validation.IsAlpha(name)
	// manipulation is remove extra spaces
	return strings.TrimSpace(name)
}

// ProcessAmount removes prefixed zeroes and puts two decimal
func ProcessAmount(amount string) string {
	validation.IsNumeric(amount)
	i, err := strconv.Atoi(amount)
	if err != nil {
		panic(err)
	}
	fl := float64(i) / 100
	return strconv.FormatFloat(fl, 'f', 2, 64)
}
