package fake

import (
	"testing"
)

func TestMemorableStrings(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := MemorableString()
		if v == "" {
			t.Errorf("MemorableString failed with lang %s", lang)
		}
	}
}
