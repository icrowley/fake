package fake

import (
	"strings"
)

func Gender() string {
	return lookup(lang, "genders", true)
}

func GenderAbbrev() string {
	g := Gender()
	if g != "" {
		return strings.ToLower(string(g[0]))
	}
	return ""
}

func Language() string {
	return lookup(lang, "languages", true)
}
