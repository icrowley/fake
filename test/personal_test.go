package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestPersonal(t *testing.T) {
	fake.SetLang("en")

	v := fake.Gender()
	if v == "" {
		t.Error("Gender failed")
	}

	v = fake.GenderAbbrev()
	if v == "" {
		t.Error("GenderAbbrev failed")
	}

	v = fake.Language()
	if v == "" {
		t.Error("Language failed")
	}
}
