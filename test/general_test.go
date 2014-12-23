package test

import (
	"testing"

	"github.com/icrowley/fake"
)

func TestGeneral(t *testing.T) {
	langs := []string{"en", "ru"}
	for _, lang := range langs {
		fake.SetLang(lang)

		v := fake.Password(4, 10, true, true, true)
		if v == "" {
			t.Errorf("Password failed with lang %s", lang)
		}

		v = fake.SimplePassword()
		if v == "" {
			t.Errorf("SimplePassword failed with lang %s", lang)
		}

		v = fake.Color()
		if v == "" {
			t.Errorf("Color failed with lang %s", lang)
		}

		v = fake.HexColor()
		if v == "" {
			t.Errorf("HexColor failed with lang %s", lang)
		}

		v = fake.ShortHexColor()
		if v == "" {
			t.Errorf("ShortHexColor failed with lang %s", lang)
		}

		v = fake.DigitsN(2)
		if v == "" {
			t.Errorf("DigitsN failed with lang %s", lang)
		}

		v = fake.Digits()
		if v == "" {
			t.Errorf("Digits failed with lang %s", lang)
		}
	}
}
