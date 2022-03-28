package city

import (
	"rebrainpkg/wordz"
)

func Digit() string {
	wordz.Prefix = "Random string number: "
	wordz.Words = []string{"one", "two", "three", "four", "five"}
	return wordz.Random()
}
