package fake

import (
	"strings"
)

// Gender generates random gender
func Gender() string {
	return f.Gender()
}

// GenderAbbrev returns first downcased letter of the random gender
func GenderAbbrev() string {
	return f.GenderAbbrev()
}

// Language generates random human language
func Language() string {
	return f.Language()
}

// Gender generates random gender
func (f *Faker) Gender() string {
	return f.lookup(f.lang, "genders", true)
}

// GenderAbbrev returns first downcased letter of the random gender
func (f *Faker) GenderAbbrev() string {
	g := Gender()
	if g != "" {
		return strings.ToLower(string(g[0]))
	}
	return ""
}

// Language generates random human f.language
func (f *Faker) Language() string {
	return f.lookup(f.lang, "languages", true)
}
