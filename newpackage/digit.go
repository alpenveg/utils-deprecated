package newpackage

import (
	"gopackages/wordz"
)

func Digit() string {
	wordz.Prefix = "Random number: "
	wordz.Words = []string{"One", "Two", "Three", "Four", "Five"}

	return wordz.Random()
}
