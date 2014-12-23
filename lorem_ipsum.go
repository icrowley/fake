package fake

import (
	"strings"
)

func Character() string {
	return lookup(lang, "characters", true)
}

func CharactersN(n int) string {
	var chars []string
	for i := 0; i < n; i++ {
		chars = append(chars, Character())
	}
	return strings.Join(chars, "")
}

func Characters() string {
	return CharactersN(r.Intn(5) + 1)
}

func Word() string {
	return lookup(lang, "words", true)
}

func WordsN(n int) string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = Word()
	}
	return strings.Join(words, " ")
}

func Words() string {
	return WordsN(r.Intn(5) + 1)
}

func Title() string {
	return strings.ToTitle(WordsN(2 + r.Intn(3)))
}

func Sentence() string {
	var words []string
	for i := 0; i < 3+r.Intn(12); i++ {
		word := Word()
		if r.Intn(5) == 0 {
			word += ","
		}
		words = append(words, Word())
	}

	sentence := strings.Join(words, " ")

	if r.Intn(8) == 0 {
		sentence += "!"
	} else {
		sentence += "."
	}

	return sentence
}

func SentencesN(n int) string {
	sentences := make([]string, n)
	for i := 0; i < n; i++ {
		sentences[i] = Sentence()
	}
	return strings.Join(sentences, " ")
}

func Sentences() string {
	return SentencesN(r.Intn(5) + 1)
}

func Paragraph() string {
	return SentencesN(r.Intn(10) + 1)
}

func ParagraphsN(n int) string {
	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, Paragraph())
	}
	return strings.Join(paragraphs, "\t")
}

func Paragraphs() string {
	return ParagraphsN(r.Intn(5) + 1)
}
