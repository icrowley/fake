package fake

import "strings"

// MemorableString generates an easy to remember string in the format AdjectiveAdjectiveAnimal
func MemorableString() string {
	return strings.Title(lookup(lang, "memorable_adjectives", true)) +
		strings.Title(lookup(lang, "memorable_adjectives", true)) +
		strings.Title(lookup(lang, "memorable_animals", true))
}
