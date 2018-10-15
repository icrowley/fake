package fake

import (
	"strings"
)

// Character generates random character in the given language
func Character() string {
	return f.Character()
}

// CharactersN generates n random characters in the given language
func CharactersN(n int) string {
	return f.CharactersN(n)
}

// Characters generates from 1 to 5 characters in the given language
func Characters() string {
	return f.Characters()
}

// Word generates random word
func Word() string {
	return f.Word()
}

// WordsN generates n random words
func WordsN(n int) string {
	return f.WordsN(n)
}

// Words generates from 1 to 5 random words
func Words() string {
	return f.Words()
}

// Title generates from 2 to 5 titleized words
func Title() string {
	return f.Title()
}

// Sentence generates random sentence
func Sentence() string {
	return f.Sentence()
}

// SentencesN generates n random sentences
func SentencesN(n int) string {
	return f.SentencesN(n)
}

// Sentences generates from 1 to 5 random sentences
func Sentences() string {
	return f.Sentences()
}

// Paragraph generates paragraph
func Paragraph() string {
	return f.Paragraph()
}

// ParagraphsN generates n paragraphs
func ParagraphsN(n int) string {
	return f.ParagraphsN(n)
}

// Paragraphs generates from 1 to 5 paragraphs
func Paragraphs() string {
	return f.Paragraphs()
}

// Character generates random character in the given f.language
func (f *Faker) Character() string {
	return f.lookup(f.lang, "characters", true)
}

// CharactersN generates n random characters in the given f.language
func (f *Faker) CharactersN(n int) string {
	var chars []string
	for i := 0; i < n; i++ {
		chars = append(chars, Character())
	}
	return strings.Join(chars, "")
}

// Characters generates from 1 to 5 characters in the given f.language
func (f *Faker) Characters() string {
	return f.CharactersN(f.r.Intn(5) + 1)
}

// Word generates random word
func (f *Faker) Word() string {
	return f.lookup(f.lang, "words", true)
}

// WordsN generates n random words
func (f *Faker) WordsN(n int) string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = Word()
	}
	return strings.Join(words, " ")
}

// Words generates from 1 to 5 random words
func (f *Faker) Words() string {
	return f.WordsN(f.r.Intn(5) + 1)
}

// Title generates from 2 to 5 titleized words
func (f *Faker) Title() string {
	return strings.ToTitle(WordsN(2 + f.r.Intn(4)))
}

// Sentence generates random sentence
func (f *Faker) Sentence() string {
	var words []string
	for i := 0; i < 3+f.r.Intn(12); i++ {
		word := Word()
		if f.r.Intn(5) == 0 {
			word += ","
		}
		words = append(words, Word())
	}

	sentence := strings.Join(words, " ")

	if f.r.Intn(8) == 0 {
		sentence += "!"
	} else {
		sentence += "."
	}

	return sentence
}

// SentencesN generates n random sentences
func (f *Faker) SentencesN(n int) string {
	sentences := make([]string, n)
	for i := 0; i < n; i++ {
		sentences[i] = Sentence()
	}
	return strings.Join(sentences, " ")
}

// Sentences generates from 1 to 5 random sentences
func (f *Faker) Sentences() string {
	return f.SentencesN(f.r.Intn(5) + 1)
}

// Paragraph generates paragraph
func (f *Faker) Paragraph() string {
	return f.SentencesN(f.r.Intn(10) + 1)
}

// ParagraphsN generates n paragraphs
func (f *Faker) ParagraphsN(n int) string {
	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, Paragraph())
	}
	return strings.Join(paragraphs, "\t")
}

// Paragraphs generates from 1 to 5 paragraphs
func (f *Faker) Paragraphs() string {
	return f.ParagraphsN(f.r.Intn(5) + 1)
}
