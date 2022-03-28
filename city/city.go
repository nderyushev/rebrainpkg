package city

import (
	"rebrainpkg/wordz"
)

func City() string {
	wordz.Prefix = "Random city: "
	wordz.Words = []string{"Moscow", "Paris", "Saint-P", "Barcelona", "New York"}
	return wordz.Random()
}
