package fake

import (
	"testing"
)

func TestFiles(t *testing.T) {
	for _, lang := range GetLangs() {
		SetLang(lang)

		v := FileName()
		if v == "" {
			t.Errorf("FileName failed with lang %s", lang)
		}

		v = FileExtension()
		if v == "" {
			t.Errorf("FileExtension failed with lang %s", lang)
		}
	}
}
