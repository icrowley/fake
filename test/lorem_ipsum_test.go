package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestLoremIpsum(t *testing.T) {
	fake.SetLang("en")

	v := fake.Character()
	if v == "" {
		t.Error("Character failed")
	}

	v = fake.CharactersN(2)
	if v == "" {
		t.Error("CharactersN failed")
	}

	v = fake.Characters()
	if v == "" {
		t.Error("Characters failed")
	}

	v = fake.Word()
	if v == "" {
		t.Error("Word failed")
	}

	v = fake.WordsN(2)
	if v == "" {
		t.Error("WordsN failed")
	}

	v = fake.Words()
	if v == "" {
		t.Error("Words failed")
	}

	v = fake.Title()
	if v == "" {
		t.Error("Title failed")
	}

	v = fake.Sentence()
	if v == "" {
		t.Error("Sentence failed")
	}

	v = fake.SentencesN(2)
	if v == "" {
		t.Error("SentencesN failed")
	}

	v = fake.Sentences()
	if v == "" {
		t.Error("Sentences failed")
	}

	v = fake.Paragraph()
	if v == "" {
		t.Error("Paragraph failed")
	}

	v = fake.ParagraphsN(2)
	if v == "" {
		t.Error("ParagraphsN failed")
	}

	v = fake.Paragraphs()
	if v == "" {
		t.Error("Paragraphs failed")
	}
}
