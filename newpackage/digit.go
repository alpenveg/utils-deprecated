package newpackage

import (
	"github.com/alpenveg/utils/wordz"
)

func Digit() string {
	wordz.Prefix = "Random number: "
	wordz.Words = []string{"One", "Two", "Three", "Four", "Five"}

	return wordz.Random()
}
